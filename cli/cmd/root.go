package cmd

import "github.com/spf13/cobra"

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

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(lockCmd)
	rootCmd.AddCommand(validateCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
