package models

import "github.com/spf13/cobra"

type Arg struct {
	Name        string
	Short       string
	Description string
	Required    bool
}

type StringArg struct {
	*Arg
	Value string
}

type StringSliceArg struct {
	*Arg
	Value []string
}

type BoolArg struct {
	*Arg
	Value bool
}

func (arg *Arg) RegisterWithCmd(cmd *cobra.Command) {
	if arg.Required {
		cmd.MarkFlagRequired(arg.Name)
	}
}

func (arg *StringArg) RegisterWithCmd(cmd *cobra.Command) {
	if arg.Short != "" {
		cmd.Flags().StringVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
	} else {
		cmd.Flags().StringVar(&arg.Value, arg.Name, arg.Value, arg.Description)
	}
	if arg.Required {
		cmd.MarkFlagRequired(arg.Name)
	}
}

func (arg *StringArg) PersistentRegisterWithCmd(cmd *cobra.Command) {
	arg.Arg.RegisterWithCmd(cmd)
	if arg.Short != "" {
		cmd.PersistentFlags().StringVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
	} else {
		cmd.PersistentFlags().StringVar(&arg.Value, arg.Name, arg.Value, arg.Description)
	}
}

func (arg *StringSliceArg) RegisterWithCmd(cmd *cobra.Command) {
	arg.Arg.RegisterWithCmd(cmd)
	if arg.Short != "" {
		cmd.Flags().StringSliceVarP(&arg.Value, arg.Name, arg.Short, []string{}, arg.Description)
	} else {
		cmd.Flags().StringSliceVar(&arg.Value, arg.Name, []string{}, arg.Description)
	}
}

func (arg *StringSliceArg) PersistentRegisterWithCmd(cmd *cobra.Command) {
	arg.Arg.RegisterWithCmd(cmd)
	if arg.Short != "" {
		cmd.PersistentFlags().StringSliceVarP(&arg.Value, arg.Name, arg.Short, []string{}, arg.Description)
	} else {
		cmd.PersistentFlags().StringSliceVar(&arg.Value, arg.Name, []string{}, arg.Description)
	}
}

func (arg *BoolArg) RegisterWithCmd(cmd *cobra.Command) {
	arg.Arg.RegisterWithCmd(cmd)
	if arg.Short != "" {
		cmd.Flags().BoolVarP(&arg.Value, arg.Name, arg.Short, false, arg.Description)
	} else {
		cmd.Flags().BoolVar(&arg.Value, arg.Name, false, arg.Description)
	}
}

func (arg *BoolArg) PersistentRegisterWithCmd(cmd *cobra.Command) {
	arg.Arg.RegisterWithCmd(cmd)
	if arg.Short != "" {
		cmd.PersistentFlags().BoolVarP(&arg.Value, arg.Name, arg.Short, false, arg.Description)
	} else {
		cmd.PersistentFlags().BoolVar(&arg.Value, arg.Name, false, arg.Description)
	}
}
