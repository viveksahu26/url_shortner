package options

import "github.com/spf13/cobra"

type ShortOptions struct {
	URL string
}

// var _ Interface = (*ShortOptions)(nil)

func (o *ShortOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.URL, "URL", "", "path of the long url")
}
