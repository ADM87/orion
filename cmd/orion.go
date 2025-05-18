package cmds

import (
	"github.com/ADM87/orion/system"
	"github.com/spf13/cobra"
)

var orionCmd = &cobra.Command{
	Use:   "orion",
	Short: "Orion CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		system.SetVerbose()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	system.Verbose.AddTo(orionCmd.PersistentFlags().BoolVarP)
}

func Execute(version string) {
	orionCmd.Version = version
	if err := orionCmd.Execute(); err != nil {
		system.Fatal(err)
	}
}
