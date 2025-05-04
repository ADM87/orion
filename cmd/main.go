package main

import (
	"github.com/ADM87/orion/cmd/op"
	"github.com/ADM87/orion/system"
	"github.com/ADM87/orion/system/logging"
	"github.com/ADM87/orion/system/types"
	"github.com/spf13/cobra"
)

var (
	VerboseArg = &types.BoolArg{
		Arg:   &types.Arg{Name: "verbose", Description: "Enable verbose logging"},
		Value: false,
	}
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
	VerboseArg.PersistentRegisterWithCmd(root)

	root.AddCommand(op.InitCmd)
}

func main() {
	if err := root.Execute(); err != nil {
		system.Fatal(err)
	}
}
