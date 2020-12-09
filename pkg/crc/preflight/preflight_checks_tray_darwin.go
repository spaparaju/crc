package preflight

import (
	"bytes"
	"fmt"
	"io/ioutil"
	goos "os"
	"path/filepath"

	"github.com/Masterminds/semver"
	"github.com/code-ready/crc/pkg/crc/constants"
	"github.com/code-ready/crc/pkg/crc/logging"
	"github.com/code-ready/crc/pkg/crc/version"
	dl "github.com/code-ready/crc/pkg/download"
	"github.com/code-ready/crc/pkg/embed"
	"github.com/code-ready/crc/pkg/extract"
	"github.com/code-ready/crc/pkg/os"
	"github.com/code-ready/crc/pkg/os/launchd"
	"github.com/pkg/errors"
	"howett.net/plist"
)

const (
	daemonAgentLabel = "crc.daemon"
	trayAgentLabel   = "crc.tray"
)

var (
	stdOutFilePathDaemon = filepath.Join(constants.CrcBaseDir, ".crcd-agent.log")
	stdOutFilePathTray   = filepath.Join(constants.CrcBaseDir, ".crct-agent.log")
)

type TrayVersion struct {
	ShortVersion string `plist:"CFBundleShortVersionString"`
}

func checkIfDaemonPlistFileExists() error {
	// crc setup can be ran from any location in the
	// users computer, and we need to update the plist
	// file with the path of the crc executable which was
	// used to run setup, to force it this check needs
	// to always fail so the fix routine is triggered

	return fmt.Errorf("Ignoring this check and triggering creation of daemon plist")
}

func fixDaemonPlistFileExists() error {
	// Try to remove the daemon agent from launchd
	// and recreate its plist
	_ = launchd.Remove(daemonAgentLabel)

	currentExecutablePath, err := goos.Executable()
	if err != nil {
		return err
	}
	daemonConfig := launchd.AgentConfig{
		Label:          daemonAgentLabel,
		ExecutablePath: currentExecutablePath,
		StdOutFilePath: stdOutFilePathDaemon,
		Args:           []string{"daemon"},
	}
	return fixPlistFileExists(daemonConfig)
}

func removeDaemonPlistFile() error {
	return launchd.RemovePlist(daemonAgentLabel)
}

func checkIfTrayPlistFileExists() error {
	if !launchd.PlistExists(trayAgentLabel) {
		return fmt.Errorf("Tray plist file does not exist")
	}
	return nil
}

func fixTrayPlistFileExists() error {
	trayConfig := launchd.AgentConfig{
		Label:          trayAgentLabel,
		ExecutablePath: constants.TrayExecutablePath,
		StdOutFilePath: stdOutFilePathTray,
	}
	return fixPlistFileExists(trayConfig)
}

func removeTrayPlistFile() error {
	return launchd.RemovePlist(trayAgentLabel)
}

func checkIfDaemonAgentRunning() error {
	if !launchd.AgentRunning(daemonAgentLabel) {
		return fmt.Errorf("crc daemon is not running")
	}
	return nil
}

func fixDaemonAgentRunning() error {
	logging.Debug("Starting daemon agent")
	if err := launchd.LoadPlist(daemonAgentLabel); err != nil {
		return err
	}
	return launchd.StartAgent(daemonAgentLabel)
}

func unLoadDaemonAgent() error {
	return launchd.UnloadPlist(daemonAgentLabel)
}

func checkIfTrayAgentRunning() error {
	if !launchd.AgentRunning(trayAgentLabel) {
		return fmt.Errorf("Tray is not running")
	}
	return nil
}

func fixTrayAgentRunning() error {
	logging.Debug("Starting tray agent")
	if err := launchd.LoadPlist(trayAgentLabel); err != nil {
		return err
	}
	return launchd.StartAgent(trayAgentLabel)
}

func unLoadTrayAgent() error {
	return launchd.UnloadPlist(trayAgentLabel)
}

func checkTrayVersion() error {
	v, err := getTrayVersion(constants.TrayAppBundlePath)
	if err != nil {
		logging.Error(err.Error())
		return err
	}
	currentVersion, err := semver.NewVersion(v)
	if err != nil {
		logging.Error(err.Error())
		return err
	}
	expectedVersion, err := semver.NewVersion(version.GetCRCMacTrayVersion())
	if err != nil {
		logging.Error(err.Error())
		return err
	}

	if expectedVersion.GreaterThan(currentVersion) {
		return fmt.Errorf("Cached version is older then latest version: %s < %s", currentVersion.String(), expectedVersion.String())
	}
	return nil
}

func fixTrayVersion() error {
	logging.Debug("Downloading/extracting tray executable")
	// get the tray app
	err := downloadOrExtractTrayApp()
	if err != nil {
		return err
	}
	return launchd.RestartAgent(trayAgentLabel)
}

func checkTrayExecutablePresent() error {
	if !os.FileExists(constants.TrayExecutablePath) {
		return fmt.Errorf("Tray executable does not exist")
	}
	return nil
}

func fixTrayExecutablePresent() error {
	logging.Debug("Downloading/extracting tray executable")
	return downloadOrExtractTrayApp()
}

func fixPlistFileExists(agentConfig launchd.AgentConfig) error {
	logging.Debugf("Creating plist for %s", agentConfig.Label)
	err := launchd.CreatePlist(agentConfig)
	if err != nil {
		return err
	}
	// load plist
	if err := launchd.LoadPlist(agentConfig.Label); err != nil {
		logging.Debug("failed while creating plist:", err.Error())
		return err
	}
	return nil
}

func downloadOrExtractTrayApp() error {
	// Extract the tray and put it in the bin directory.
	tmpArchivePath, err := ioutil.TempDir("", "crc")
	if err != nil {
		logging.Error("Failed creating temporary directory for extracting tray")
		return err
	}
	defer func() {
		_ = goos.RemoveAll(tmpArchivePath)
	}()

	logging.Debug("Trying to extract tray from crc executable")
	err = embed.Extract(filepath.Base(constants.GetCRCMacTrayDownloadURL()), tmpArchivePath)
	if err != nil {
		logging.Debug("Could not extract tray from crc executable", err)
		logging.Debug("Downloading crc tray")
		_, err = dl.Download(constants.GetCRCMacTrayDownloadURL(), tmpArchivePath, 0600)
		if err != nil {
			return err
		}
	}
	archivePath := filepath.Join(tmpArchivePath, filepath.Base(constants.GetCRCMacTrayDownloadURL()))
	outputPath := constants.CrcBinDir
	err = goos.MkdirAll(outputPath, 0750)
	if err != nil {
		return errors.Wrap(err, "Cannot create the target directory.")
	}
	_, err = extract.Uncompress(archivePath, outputPath, false)
	if err != nil {
		return errors.Wrapf(err, "Cannot uncompress '%s'", archivePath)
	}
	return nil
}

func getTrayVersion(trayAppPath string) (string, error) {
	var version TrayVersion
	f, err := ioutil.ReadFile(filepath.Join(trayAppPath, "Contents", "Info.plist")) // #nosec G304
	if err != nil {
		return "", err
	}
	decoder := plist.NewDecoder(bytes.NewReader(f))
	err = decoder.Decode(&version)
	if err != nil {
		return "", err
	}

	return version.ShortVersion, nil
}
