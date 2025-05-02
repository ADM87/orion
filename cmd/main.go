package main

import (
	"errors"

	"github.com/ADM87/orion/system"
	"github.com/ADM87/orion/system/logging"
	"github.com/spf13/cobra"
)

var version = "0.0.0-unreleased"

var root = &cobra.Command{
	Use:   "orion",
	Short: "Orion CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no command specified. Use --help for more information")
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	root.Version = version
}

func main() {
	system.LogLvl(logging.LogLvlDebug | logging.LogLvlInfo | logging.LogLvlWarn | logging.LogLvlError | logging.LogLvlFatal)
	if err := root.Execute(); err != nil {
		system.Debug(err)
	}
}
