package cmds

import (
	"github.com/ADM87/orion/system"
	"github.com/ADM87/orion/system/logging"
	"github.com/spf13/cobra"
)

var orionCmd = &cobra.Command{
	Use: "orion",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		verbose := cmd.Flag("verbose")
		if verbose != nil && verbose.Value.String() == "true" {
			system.LogLvl(logging.LogLvlAll)
		}
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute(version string) {
	orionCmd.Version = version
	if err := orionCmd.Execute(); err != nil {
		system.Fatal(err)
	}
}
