package ops

import (
	"github.com/ADM87/orion/system/types"
	"github.com/spf13/cobra"
)

var (
	PathArg = &types.StringArg{
		Arg: &types.Arg{Name: "path", Description: "Path to initialize the new resource in"},
	}
	SpecArg = &types.StringArg{
		Arg:   &types.Arg{Name: "spec", Short: "s", Description: "Path to the resource spec file (spec.yaml)"},
		Value: "spec.yaml",
	}
	NoPromptArg = &types.BoolArg{
		Arg: &types.Arg{Name: "no-prompt", Short: "y", Description: "Skip confirmation prompt"},
	}
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	PathArg.PersistentRegisterWithCmd(InitCmd)
	NoPromptArg.PersistentRegisterWithCmd(InitCmd)

	SpecArg.RegisterWithCmd(InitCmd)
}
