package main

import (
	"github.com/ADM87/orion/models"
	"github.com/ADM87/orion/system"
	"github.com/ADM87/orion/system/logging"
	"github.com/spf13/cobra"
)

var version = "0.0.0-unreleased"
var root = &cobra.Command{
	Use:     "orion",
	Short:   "Orion CLI",
	Version: version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		verbose := cmd.Flag("verbose")
		if verbose != nil && verbose.Value.String() == "true" {
			system.LogLvl(logging.LogLvlAll)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	models.VerboseArg.PersistentRegisterWithCmd(root)
}

func main() {
	if err := root.Execute(); err != nil {
		system.Fatal(err)
	}
}
