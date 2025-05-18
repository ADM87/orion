package cmds

import (
	cmdargs "github.com/ADM87/orion/cmd/args"
	"github.com/ADM87/orion/system"
	"github.com/spf13/cobra"
)

var orionCmd = &cobra.Command{
	Use: "orion",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		system.SetVerbose(cmdargs.Verbose.GetValue())
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cmdargs.Verbose.AddPersistentFlag(orionCmd)
}

func Execute(version string) {
	orionCmd.Version = version
	if err := orionCmd.Execute(); err != nil {
		system.Fatal(err)
	}
}
