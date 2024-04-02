package server

import (
	"fmt"

	pb "github.com/iliyankg/colab-shield/protos"
)

func buildRedisKeyForFile(projectId string, fileId string) string {
	return fmt.Sprintf("project:%s:file:%s", projectId, fileId)
}

func keysFromFileClaimRequests(target *[]string, projectId string, files []*pb.FileClaimRequest) {
	for _, file := range files {
		*target = append(*target, buildRedisKeyForFile(projectId, file.FileId))
	}
}
