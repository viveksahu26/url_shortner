package cli

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "url_shortner",
		Short:             "tool to convert long url into short",
		DisableAutoGenTag: true,
		SilenceUsage:      true, // Don't show usage on errors
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	// ro.AddFlags(cmd)

	// templates.SetCustomUsageFunc(cmd)

	// Add sub-commands.

	cmd.AddCommand(Short())

	return cmd
}
