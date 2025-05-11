package op

import (
	"fmt"

	"github.com/ADM87/orion/models"
	"github.com/ADM87/orion/system"
	"github.com/ADM87/orion/system/templating"
	"github.com/spf13/cobra"
)

var (
	SpecArg = &models.StringArg{
		Arg:   &models.Arg{Name: "spec", Short: "s", Description: fmt.Sprintf("Path to the resource spec file (%s)", templating.TemplatingSpecFileName)},
		Value: "spec.yaml",
	}
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := system.MakeAbsolute(SpecArg.Value)
		if err != nil {
			return err
		}

		return nil
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	models.PathArg.PersistentRegisterWithCmd(InitCmd)
	models.NoPromptArg.PersistentRegisterWithCmd(InitCmd)

	SpecArg.RegisterWithCmd(InitCmd)
}
