package op

import (
	"fmt"

	"github.com/ADM87/orion/system"
	"github.com/ADM87/orion/system/templating"
	"github.com/ADM87/orion/system/types"
	"github.com/spf13/cobra"
)

var (
	PathArg = &types.StringArg{
		Arg: &types.Arg{Name: "path", Description: "Path to initialize the new resource in"},
	}
	SpecArg = &types.StringArg{
		Arg:   &types.Arg{Name: "spec", Short: "s", Description: fmt.Sprintf("Path to the resource spec file (%s)", templating.TemplatingSpecFileName)},
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
		spec, err := system.MakeAbsolute(SpecArg.Value)
		if err != nil {
			return err
		}

		if _, err := system.LoadTemplateSpec(spec); err != nil {
			return err
		}
		return nil
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	PathArg.PersistentRegisterWithCmd(InitCmd)
	NoPromptArg.PersistentRegisterWithCmd(InitCmd)

	SpecArg.RegisterWithCmd(InitCmd)
}
