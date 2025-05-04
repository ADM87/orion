package ops

import (
	"github.com/ADM87/orion/system/types"
	"github.com/spf13/cobra"
)

var (
	PathArg = &types.StringArg{
		Arg: &types.Arg{Name: "path"},
	}
	NoPromptArg = &types.BoolArg{
		Arg: &types.Arg{Name: "no-prompt"},
	}
)

var NewCmd = &cobra.Command{
	Use: "new",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	PathArg.PersistentRegisterWithCmd(NewCmd)
	NoPromptArg.PersistentRegisterWithCmd(NewCmd)
}
