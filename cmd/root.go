package cmd

import (
	"context"
	"fmt"

	config "github.com/open-cloud-initiative/cli/internal/cfg"
	"github.com/open-cloud-initiative/cli/pkg/extensions"

	"github.com/spf13/cobra"
)

var cfg = config.New()

const versionFmt = "%s (%s %s)"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var mgr extensions.Manager

func Init() error {
	ctx := context.Background()

	err := cfg.InitDefaultConfig()
	if err != nil {
		return err
	}

	mgr = extensions.NewManager()
	err = mgr.Scan(extensions.DataDir())
	if err != nil {
		return err
	}

	RootCmd.AddCommand(ExtCmd)

	for _, ext := range mgr.ListExtensions() {
		RootCmd.AddCommand(ext.Cmd())
	}

	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Verbose, "verbose", "v", cfg.Flags.Verbose, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Dry, "dry", "d", cfg.Flags.Dry, "dry run")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Root, "root", "r", cfg.Flags.Root, "run as root")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Force, "force", "f", cfg.Flags.Force, "force init")

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true

	err = RootCmd.ExecuteContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

var RootCmd = &cobra.Command{
	Use:   "ocictl",
	Short: "ocictl",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context(), args...)
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
}

func runRoot(_ context.Context, _ ...string) error {
	return nil
}
