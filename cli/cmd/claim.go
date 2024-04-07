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
)

func init() {
	claimFilesCmd.Flags().StringArrayVarP(&filesToClaim, "file", "f", []string{}, "files to lock")
	claimFilesCmd.MarkFlagRequired("file")
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
		payload, err := newClaimFilesRequest(filesToClaim, hashes, pb.ClaimMode_EXCLUSIVE)
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

		if response.Status != pb.Status_OK {
			log.Fatal().Msg("status not OK")
		}
	},
}

func newClaimFilesRequest(files []string, hashes []string, claimMode pb.ClaimMode) (*pb.ClaimFilesRequest, error) {
	if len(files) != len(hashes) {
		return nil, ErrFileToHashMissmatch
	}

	claimFileInfos := make([]*pb.ClaimFileInfo, 0, len(filesToClaim))
	for i, file := range files {
		claimFileInfos = append(claimFileInfos, &pb.ClaimFileInfo{
			FileId:    file,
			FileHash:  hashes[i],
			ClaimMode: claimMode,
		})
	}

	return &pb.ClaimFilesRequest{
		BranchName: gitBranch,
		Files:      claimFileInfos,
	}, nil
}
