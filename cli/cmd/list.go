package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all locked files.",
	Long:  `Lists all locked files and all respective metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("List called")
	},
}
