package ssh

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/code-ready/crc/pkg/crc/constants"
	"github.com/code-ready/crc/pkg/crc/logging"
	crcos "github.com/code-ready/crc/pkg/os"
)

type Runner struct {
	ip             string
	port           int
	privateSSHKeys []string
}

func CreateRunner(ip string, port int, privateKeys ...string) *Runner {
	return &Runner{
		ip:             ip,
		port:           port,
		privateSSHKeys: privateKeys,
	}
}

// Create a host using the driver's config
func (runner *Runner) Run(command string) (string, error) {
	return runner.runSSHCommandFromDriver(command, false)
}

func (runner *Runner) RunPrivate(command string) (string, error) {
	return runner.runSSHCommandFromDriver(command, true)
}

func (runner *Runner) CopyData(data []byte, destFilename string, mode os.FileMode) error {
	logging.Debugf("Creating %s with permissions 0%o in the CRC VM", destFilename, mode)
	base64Data := base64.StdEncoding.EncodeToString(data)
	command := fmt.Sprintf("sudo install -m 0%o /dev/null %s && cat <<EOF | base64 --decode | sudo tee %s\n%s\nEOF", mode, destFilename, destFilename, base64Data)
	_, err := runner.RunPrivate(command)

	return err
}

func (runner *Runner) CopyFile(srcFilename string, destFilename string, mode os.FileMode) error {
	data, err := ioutil.ReadFile(srcFilename)
	if err != nil {
		return err
	}
	return runner.CopyData(data, destFilename, mode)
}

func (runner *Runner) runSSHCommandFromDriver(command string, runPrivate bool) (string, error) {
	var availableKeys []string
	for _, privateKey := range runner.privateSSHKeys {
		if _, err := os.Stat(privateKey); err == nil {
			availableKeys = append(availableKeys, privateKey)
		}
	}

	client, err := NewClient(constants.DefaultSSHUser, runner.ip, runner.port, &Auth{
		Keys: availableKeys,
	})
	if err != nil {
		return "", err
	}

	if runPrivate {
		logging.Debugf("About to run SSH command with hidden output")
	} else {
		logging.Debugf("About to run SSH command:\n%s", command)
	}

	output, err := client.Output(command)
	if runPrivate {
		if err != nil {
			logging.Debugf("SSH command failed")
		} else {
			logging.Debugf("SSH command succeeded")
			logging.Debugf("SSH command results: , output: %s", output)

		}
	} else {
		logging.Debugf("SSH command results: err: %v, output: %s", err, output)
	}

	if err != nil {
		return "", fmt.Errorf(`ssh command error:
command : %s
err     : %v
output  : %s`, command, err, output)
	}

	return output, nil
}

type remoteCommandRunner struct {
	sshRunner *Runner
}

func (cmdRunner *remoteCommandRunner) Run(cmd string, args ...string) (string, string, error) {
	commandline := fmt.Sprintf("%s %s", cmd, strings.Join(args, " "))
	out, err := cmdRunner.sshRunner.Run(commandline)
	return out, "", err
}

func (cmdRunner *remoteCommandRunner) RunPrivate(cmd string, args ...string) (string, string, error) {
	commandline := fmt.Sprintf("%s %s", cmd, strings.Join(args, " "))
	out, err := cmdRunner.sshRunner.RunPrivate(commandline)
	return out, "", err
}

func (cmdRunner *remoteCommandRunner) RunPrivileged(reason string, cmdAndArgs ...string) (string, string, error) {
	commandline := fmt.Sprintf("sudo %s", strings.Join(cmdAndArgs, " "))

	out, err := cmdRunner.sshRunner.Run(commandline)

	return out, "", err
}

func NewRemoteCommandRunner(runner *Runner) crcos.CommandRunner {
	return &remoteCommandRunner{
		sshRunner: runner,
	}
}
