package systemargs

import (
	"fmt"

	"github.com/spf13/cobra"
)

type (
	CmdArgReceiver[T any] func(value *T, name, short string, defaultValue T, description string)

	CmdArg[T any] interface {
		GetValue() T

		GetName() string
		GetShort() string
		GetDescription() string

		AddTo(receiver CmdArgReceiver[T])

		Validate(cmd *cobra.Command) error
	}

	cmdarg[T any] struct {
		name        string
		short       string
		description string
		required    bool
		value       T
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

func (c *cmdarg[T]) AddTo(receiver CmdArgReceiver[T]) {
	receiver(&c.value, c.name, c.short, c.value, c.description)
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

func NewStringArg(name, short, description string, defaultValue string, required bool) CmdArg[string] {
	return &cmdarg[string]{
		name:        name,
		short:       short,
		description: description,
		required:    required,
		value:       defaultValue,
	}
}

// ======================================================================
// StringSliceArg
// ======================================================================

func NewStringSliceArg(name, short, description string, defaultValue []string, required bool) CmdArg[[]string] {
	return &cmdarg[[]string]{
		name:        name,
		short:       short,
		description: description,
		required:    required,
		value:       defaultValue,
	}
}

// ======================================================================
// BoolArg
// ======================================================================

func NewBoolArg(name, short, description string, required bool) CmdArg[bool] {
	return &cmdarg[bool]{
		name:        name,
		short:       short,
		description: description,
		required:    required,
		value:       false,
	}
}
