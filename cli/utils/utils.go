package utils

import (
	"errors"

	pb "github.com/iliyankg/colab-shield/protos"
)

var (
	ErrFileToHashMissmatch = errors.New("files and hashes must be of the same length")
)

func BuildFileClaimRequests(target *[]*pb.ClaimFileInfo, files []string, hashes []string, claimMode pb.ClaimMode) error {
	if len(files) != len(hashes) {
		return ErrFileToHashMissmatch
	}

	for i, file := range files {
		*target = append(*target, &pb.ClaimFileInfo{
			FileId:    file,
			FileHash:  hashes[i],
			ClaimMode: claimMode,
		})
	}

	return nil
}
