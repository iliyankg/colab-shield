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
)

var (
	filesToLock []string
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock file(s) for editing",
	Long:  `Lock file(s) for editing`,
	Run: func(cmd *cobra.Command, args []string) {
		hashes, err := gitutils.GetGitBlobHashes(filesToLock)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get git hashes")
		}

		log.Info().Msgf("Git hash for files %s: %s", filesToLock, hashes)

		mappedFiles := utils.BuildFileToHashMap(filesToLock, hashes)
		if mappedFiles == nil {
			log.Fatal().Msg("Failed to map files to hash")
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		conn, client := client.NewColabShieldClient(ctx, serverAddress)
		defer conn.Close()

		payload := &pb.LockRequest{
			ProjectId:  gitBranch,
			UserId:     gitUser,
			BranchName: gitBranch,
			Files:      mappedFiles,
		}

		response, err := client.Lock(ctx, payload)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to lock files")
		}

		if response.Status != pb.Status_OK {
			log.Fatal().Msg("status not OK")
		}
	},
}

func init() {
	lockCmd.Flags().StringArrayVarP(&filesToLock, "file", "f", []string{}, "files to lock")
	lockCmd.MarkFlagRequired("file")
}
