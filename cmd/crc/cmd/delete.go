package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/code-ready/crc/pkg/crc/constants"
	"github.com/code-ready/crc/pkg/crc/exit"
	"github.com/code-ready/crc/pkg/crc/input"
	"github.com/code-ready/crc/pkg/crc/machine"
	"github.com/spf13/cobra"
)

var clearCache bool

func init() {
	deleteCmd.Flags().BoolVarP(&clearCache, "clear-cache", "", false,
		fmt.Sprintf("Clear the OpenShift cluster cache at: %s", constants.MachineCacheDir))
	addOutputFormatFlag(deleteCmd)
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the OpenShift cluster",
	Long:  "Delete the OpenShift cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runDelete(os.Stdout, newMachine(), clearCache, constants.MachineCacheDir, outputFormat != jsonFormat, globalForce, outputFormat); err != nil {
			exit.WithMessage(1, err.Error())
		}
	},
}

func deleteMachine(client machine.Client, clearCache bool, cacheDir string, interactive, force bool) (bool, error) {
	if !interactive && !force {
		return false, errors.New("non-interactive deletion requires --force")
	}

	if clearCache {
		yes := input.PromptUserForYesOrNo("Do you want to delete the OpenShift cluster cache", force)
		if yes {
			_ = os.RemoveAll(cacheDir)
		}
	}

	if err := checkIfMachineMissing(client); err != nil {
		return false, err
	}

	yes := input.PromptUserForYesOrNo("Do you want to delete the OpenShift cluster", force)
	if yes {
		return true, client.Delete()
	}
	return false, nil
}

func runDelete(writer io.Writer, client machine.Client, clearCache bool, cacheDir string, interactive, force bool, outputFormat string) error {
	machineDeleted, err := deleteMachine(client, clearCache, cacheDir, interactive, force)
	return render(&deleteResult{
		Success:        err == nil,
		Error:          errorMessage(err),
		machineDeleted: machineDeleted,
	}, writer, outputFormat)
}

type deleteResult struct {
	Success        bool   `json:"success"`
	Error          string `json:"error,omitempty"`
	machineDeleted bool
}

func (s *deleteResult) prettyPrintTo(writer io.Writer) error {
	if s.Error != "" {
		return errors.New(s.Error)
	}
	var err error
	if s.machineDeleted {
		_, err = fmt.Fprintln(writer, "Deleted the OpenShift cluster")
	}
	return err
}
