package cli

import (
	"github.com/spf13/cobra"
	"github.com/viveksahu26/url_shortner/cmd/url_shortner/cli/options"
	"github.com/viveksahu26/url_shortner/cmd/url_shortner/cli/short"
)

func Short() *cobra.Command {
	o := &options.ShortOptions{}

	cmd := &cobra.Command{
		Use:   "short",
		Short: "short the provided long URL",
		Long:  `Convert the provided long URL into Short URL`,
		Example: `url_shortner short --url=<long URL>
		
		# convert the long url into short url
		url_shortner short --url="http://google.com/1346461234567890123456789/get/viveksahu266"`,

		Args: cobra.MinimumNArgs(0),
		// PersistentPreRun: options.BindViper,
		RunE: func(cmd *cobra.Command, args []string) error {
			return short.GenerateShortURL(cmd.Context(), o.URL, args)
		},
	}
	o.AddFlags(cmd)
	return cmd
}
