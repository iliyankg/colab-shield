package cmd

import (
	pb "github.com/iliyankg/colab-shield/protos"

	"github.com/iliyankg/colab-shield/cli/client"
	"github.com/iliyankg/colab-shield/cli/gitutils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	updateCmd.Flags().StringArrayVarP(&files, "file", "f", []string{}, "files to lock")
	updateCmd.MarkFlagRequired("file")
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update claimed files with changes",
	Long:  `Update claimed files with changes`,
	Run: func(cmd *cobra.Command, args []string) {
		hashes, err := gitutils.GetGitBlobHashes(&log.Logger, files)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get git hashes")
		}

		headHashes, err := gitutils.GetGitBlobHEADHashes(&log.Logger, files)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get git HEAD hashes")
		}

		payload, err := newUpdateFilesRequest(files, hashes, headHashes)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to map files to hash")
		}

		ctx, cancel := buildContext(gitRepo, gitUser)
		defer cancel()
		conn, client := client.NewColabShieldClient(ctx, serverAddress)
		defer conn.Close()

		response, err := client.Update(ctx, payload)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to update files")
		}

		if response.Status != pb.Status_OK {
			log.Fatal().Msg("Failed to update files")
		}
	},
}

func newUpdateFilesRequest(files []string, hashes []string, headHashes []string) (*pb.UpdateFilesRequest, error) {
	if len(files) != len(hashes) || len(files) != len(headHashes) {
		return nil, ErrFileToHashMissmatch
	}

	updateFileInfos := make([]*pb.UpdateFileInfo, 0, len(files))
	for i, file := range files {
		updateFileInfos = append(updateFileInfos, &pb.UpdateFileInfo{
			FileId:   file,
			FileHash: hashes[i],
			OldHash:  headHashes[i],
		})
	}

	return &pb.UpdateFilesRequest{
		BranchName: gitBranch,
		Files:      updateFileInfos,
	}, nil
}
