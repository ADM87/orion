package unity

import "github.com/spf13/cobra"

var NewPackageCmd = &cobra.Command{
	Use:   "unity-package",
	Short: "Create a new Unity package",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {

}
