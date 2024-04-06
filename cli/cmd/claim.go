package cmd

import (
	"context"
	"time"

	"github.com/iliyankg/colab-shield/cli/client"
	"github.com/iliyankg/colab-shield/cli/gitutils"
	"github.com/iliyankg/colab-shield/cli/utils"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var (
	filesToClaim []string
)

var claimFilesCmd = &cobra.Command{
	Use:   "claim",
	Short: "Claim file(s) for editing",
	Long:  `Claim file(s) for editing`,
	Run: func(cmd *cobra.Command, args []string) {
		hashes, err := gitutils.GetGitBlobHashes(filesToClaim)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get git hashes")
		}

		log.Info().Msgf("Git hash for files %s: %s", filesToClaim, hashes)
		claimFileInfos := make([]*pb.ClaimFileInfo, 0, len(filesToClaim))
		err = utils.BuildFileClaimRequests(&claimFileInfos, filesToClaim, hashes, pb.ClaimMode_EXCLUSIVE)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to map files to hash")
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		conn, client := client.NewColabShieldClient(ctx, serverAddress)
		defer conn.Close()

		mdCtx := metadata.Pairs(
			"projectId", gitRepo,
			"userId", gitUser,
		)
		ctx = metadata.NewOutgoingContext(ctx, mdCtx)

		payload := &pb.ClaimFilesRequest{
			BranchName: gitBranch,
			Files:      claimFileInfos,
		}

		response, err := client.Claim(ctx, payload)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to lock files")
		}

		if response.Status != pb.Status_OK {
			log.Fatal().Msg("status not OK")
		}
	},
}

func init() {
	claimFilesCmd.Flags().StringArrayVarP(&filesToClaim, "file", "f", []string{}, "files to lock")
	claimFilesCmd.MarkFlagRequired("file")
}
