package cmd

import (
	"context"
	"log"

	"github.com/katallaxie/pkg/slices"
	"github.com/open-cloud-initiative/cli/pkg/extensions"
	"github.com/spf13/cobra"
)

func init() {
	ExtCmd.AddCommand(ExtListCmd)
}

var ExtCmd = &cobra.Command{
	Use:   "extension",
	Short: "Manage extensions",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runExtension(cmd.Context())
	},
}

func runExtension(_ context.Context) error {
	return nil
}

var ExtListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all extensions",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runExtList(cmd.Context())
	},
}

func runExtList(_ context.Context) error {
	slices.ForEach(func(ext extensions.Extension, _ int) {
		log.Print(ext.Name())
	}, mgr.ListExtensions()...)

	return nil
}
