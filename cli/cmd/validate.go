package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	filesToValidate []string
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates if the files listed can be moodified by this user in this branch.",
	Long:  `Validates if the files listed can be moodified by this user in this branch.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal().Msg("not implemented")
	},
}

func init() {
	validateCmd.Flags().StringArrayVarP(&filesToValidate, "file", "f", []string{}, "files to validate")
	validateCmd.MarkFlagRequired("file")
}
