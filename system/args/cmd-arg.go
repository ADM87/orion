package systemargs

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CmdArg[T any] interface {
	GetValue() T

	GetName() string
	GetShort() string
	GetDescription() string

	AddFlag(cmd *cobra.Command)
	AddPersistentFlag(cmd *cobra.Command)

	Validate(cmd *cobra.Command) error
}

type cmdarg[T any] struct {
	name        string
	short       string
	description string
	required    bool
	value       T
}

type (
	stringarg struct {
		*cmdarg[string]
	}
	boolarg struct {
		*cmdarg[bool]
	}
)

// ======================================================================
// CmdArg
// ======================================================================

func (c *cmdarg[T]) GetName() string {
	return c.name
}

func (c *cmdarg[T]) GetShort() string {
	return c.short
}

func (c *cmdarg[T]) GetDescription() string {
	return c.description
}

func (c *cmdarg[T]) GetValue() T {
	return c.value
}

func (c *cmdarg[T]) AddFlag(cmd *cobra.Command) {
	if c.required {
		cmd.MarkFlagRequired(c.name)
	}
}

func (c *cmdarg[T]) AddPersistentFlag(cmd *cobra.Command) {
	if c.required {
		cmd.MarkPersistentFlagRequired(c.name)
	}
}

func (c *cmdarg[T]) Validate(cmd *cobra.Command) error {
	if c.required && cmd.Flags().Lookup(c.name).Value.String() == "" {
		return fmt.Errorf("required flag %s is not set", c.name)
	}
	return nil
}

// ======================================================================
// StringArg
// ======================================================================

func NewStringArg(name, short, description string, required bool) CmdArg[string] {
	return &stringarg{
		cmdarg: &cmdarg[string]{
			name:        name,
			short:       short,
			description: description,
			required:    required,
		},
	}
}

func (c *stringarg) AddFlag(cmd *cobra.Command) {
	c.cmdarg.AddFlag(cmd)
	cmd.Flags().StringVarP(&c.value, c.name, c.short, c.value, c.description)
}

func (c *stringarg) AddPersistentFlag(cmd *cobra.Command) {
	c.cmdarg.AddPersistentFlag(cmd)
	cmd.PersistentFlags().StringVarP(&c.value, c.name, c.short, c.value, c.description)
}

// ======================================================================
// BoolArg
// ======================================================================

func NewBoolArg(name, short, description string, required bool) CmdArg[bool] {
	return &boolarg{
		cmdarg: &cmdarg[bool]{
			name:        name,
			short:       short,
			description: description,
			required:    required,
		},
	}
}

func (c *boolarg) AddFlag(cmd *cobra.Command) {
	c.cmdarg.AddFlag(cmd)
	cmd.Flags().BoolVarP(&c.value, c.name, c.short, c.value, c.description)
}

func (c *boolarg) AddPersistentFlag(cmd *cobra.Command) {
	c.cmdarg.AddPersistentFlag(cmd)
	cmd.PersistentFlags().BoolVarP(&c.value, c.name, c.short, c.value, c.description)
}
