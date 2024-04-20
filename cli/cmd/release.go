package cmd

import (
	"github.com/iliyankg/colab-shield/cli/client"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	filesToRelease []string
)

func init() {
	claimFilesCmd.Flags().StringArrayVarP(&filesToRelease, "file", "f", []string{}, "files to lock")
	claimFilesCmd.MarkFlagRequired("file")
}

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release file(s) previously claimed.",
	Long:  `Release file(s) previously claimed.`,
	Run: func(cmd *cobra.Command, args []string) {
		payload := newReleaseFilesRequest(filesToRelease)

		ctx, cancel := buildContext(gitRepo, gitUser)
		defer cancel()
		conn, client := client.NewColabShieldClient(serverAddress)
		defer conn.Close()

		response, err := client.Release(ctx, payload)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to release files")
		}

		for _, file := range response.RejectedFiles {
			log.Info().Msgf("Rejected - %s - %s", file.FileId, file.RejectReason.String())
		}

		if response.Status != protos.Status_OK {
			log.Fatal().Msg("Failed to release files")
		}
	},
}

// newReleaseFilesRequest creates a new ReleaseFilesRequest from the given files
func newReleaseFilesRequest(filesToRelease []string) *protos.ReleaseFilesRequest {
	return &protos.ReleaseFilesRequest{
		BranchName: gitBranch,
		FileIds:    filesToRelease,
	}
}