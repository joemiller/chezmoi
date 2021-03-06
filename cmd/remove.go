package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/twpayne/go-vfs"
)

var removeCommand = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Remove a file or directory",
	RunE:    makeRunE(runRemoveCommandE),
}

func init() {
	rootCommand.AddCommand(removeCommand)
}

func runRemoveCommandE(fs vfs.FS, command *cobra.Command, args []string) error {
	targetState, err := config.getTargetState(fs)
	if err != nil {
		return err
	}
	sourceNames, err := config.getSourceNames(targetState, args)
	if err != nil {
		return err
	}
	actuator := config.getDefaultActuator(fs)
	for i, targetFileName := range args {
		if err := actuator.RemoveAll(filepath.Join(config.TargetDir, targetFileName)); err != nil && !os.IsNotExist(err) {
			return err
		}
		if err := actuator.RemoveAll(filepath.Join(config.SourceDir, sourceNames[i])); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
