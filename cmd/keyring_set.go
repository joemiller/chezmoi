package cmd

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/twpayne/go-vfs"
	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/ssh/terminal"
)

var keyringSetCommand = &cobra.Command{
	Use:   "set",
	Args:  cobra.NoArgs,
	Short: "Set a password in keyring",
	RunE:  makeRunE(config.runKeyringSetCommand),
}

func init() {
	keyringCommand.AddCommand(keyringSetCommand)

	persistentFlags := keyringSetCommand.PersistentFlags()
	persistentFlags.StringVar(&config.Keyring.Password, "password", "", "password")
}

func (c *Config) runKeyringSetCommand(fs vfs.FS, cmd *cobra.Command, args []string) error {
	passwordString := c.Keyring.Password
	if passwordString == "" {
		fmt.Print("Password: ")
		password, err := terminal.ReadPassword(syscall.Stdin)
		if err != nil {
			return err
		}
		passwordString = string(password)
	}
	return keyring.Set(c.Keyring.Service, c.Keyring.User, passwordString)
}
