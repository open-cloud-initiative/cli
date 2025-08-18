package main

import (
	"github.com/open-cloud-initiative/cli/pkg/extensions"

	"github.com/spf13/cobra"
)

var _ extensions.Extension = (*HelloExtension)(nil)

type HelloExtension struct{}

// Name implements Extension.Name.
func (e *HelloExtension) Name() string {
	return "hello"
}

// Path implements Extension.Path.
func (e *HelloExtension) Path() string {
	return extensions.Unknown
}

// Version implements Extension.Version.
func (e *HelloExtension) Version() string {
	return extensions.Unknown
}

// Owner implements Extension.Owner.
func (e *HelloExtension) Owner() string {
	return "oci"
}

// Cmd implements Extension.Cmd.
func (e *HelloExtension) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "Hello extension",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Println("Hello world!")
		},
	}
}

var Extension = HelloExtension{}
