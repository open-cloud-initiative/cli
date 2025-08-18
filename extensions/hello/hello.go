package hello

import (
	"github.com/open-cloud-initiative/cli/pkg/extensions"
	"github.com/spf13/cobra"
)

var _ extensions.Extension = &helloExtension{}

type helloExtension struct{}

// Name implements Extension.Name.
func (e *helloExtension) Name() string {
	return "hello"
}

// Path implements Extension.Path.
func (e *helloExtension) Path() string {
	return extensions.Unknown
}

// Version implements Extension.Version.
func (e *helloExtension) Version() string {
	return extensions.Unknown
}

// Owner implements Extension.Owner.
func (e *helloExtension) Owner() string {
	return "oci"
}

// Cmd implements Extension.Cmd.
func (e *helloExtension) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "Hello extension",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Hello world!")
		},
	}
}

var Extension = &helloExtension{}
