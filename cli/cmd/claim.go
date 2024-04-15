package cmd

import (
	"github.com/iliyankg/colab-shield/cli/client"
	"github.com/iliyankg/colab-shield/cli/gitutils"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	filesToClaim []string
	softClaim    bool
)

func init() {
	claimFilesCmd.Flags().StringArrayVarP(&filesToClaim, "file", "f", []string{}, "files to lock")
	claimFilesCmd.MarkFlagRequired("file")

	claimFilesCmd.Flags().BoolVarP(&softClaim, "soft-claim", "s", false, "Soft claim only exposes any files that may get rejected if any. Nothing is saved to the DB")
}

var claimFilesCmd = &cobra.Command{
	Use:   "claim",
	Short: "Claim file(s) for editing",
	Long:  `Claim file(s) for editing`,
	Run: func(cmd *cobra.Command, args []string) {
		hashes, err := gitutils.GetGitBlobHEADHashes(&log.Logger, filesToClaim)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get git hashes")
		}

		log.Info().Msgf("Git hash for files %s: %s", filesToClaim, hashes)

		// TODO: Implement proper claim mode functionality
		payload, err := newClaimFilesRequest(filesToClaim, hashes, pb.ClaimMode_EXCLUSIVE, softClaim)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to map files to hash")
		}

		ctx, cancel := buildContext(gitRepo, gitUser)
		defer cancel()
		conn, client := client.NewColabShieldClient(ctx, serverAddress)
		defer conn.Close()

		response, err := client.Claim(ctx, payload)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to lock files")
		}

		for _, file := range response.RejectedFiles {
			log.Warn().Msgf("Rejected - %s - %s", file.FileId, file.RejectReason.String())
		}

		if response.Status != pb.Status_OK {
			log.Fatal().Msg("status not OK")
		}
	},
}
