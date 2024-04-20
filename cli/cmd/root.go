package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	serverAddress string
	gitRepo       string
	gitUser       string
	gitBranch     string
	rootCmd       = &cobra.Command{
		Use:   "colab-shield",
		Short: "A CLI tool for colaborative work with hard to merge files.",
		Long:  `A CLI tool for colaborative work with hard to merge files. It does this by providing an interface to a backend server which tracks file changes and versions.`,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&serverAddress, "serverAddress", "s", "localhost:8080", "address of the server")
	rootCmd.MarkFlagRequired("serverAddress")

	rootCmd.PersistentFlags().StringVarP(&gitRepo, "gitRepo", "r", "", "git repo name")
	rootCmd.MarkFlagRequired("gitRepo")

	rootCmd.PersistentFlags().StringVarP(&gitUser, "gitUser", "u", "", "git user name")
	rootCmd.MarkFlagRequired("gitUser")

	rootCmd.PersistentFlags().StringVarP(&gitBranch, "gitBranch", "b", "", "git branch")
	rootCmd.MarkFlagRequired("gitBranch")

	rootCmd.AddCommand(release)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(initProjectCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(claimFilesCmd)
}

// Execute executes the root command.
func Execute() error {
	// Ensure context is root of a git repository
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		log.Error().Msg("Make sure you are in the root of a git repository!")
		log.Fatal().Msg(".git folder does not exist")
	}

	return rootCmd.Execute()
}
