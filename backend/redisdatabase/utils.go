package redisdatabase

import (
	"fmt"

	"github.com/iliyankg/colab-shield/backend/domain"
)

// buildRedisKeyForFile builds the Redis key for the given project and file IDs
func buildRedisKeyForFile(projectId string, fileId string) string {
	return fmt.Sprintf("project:%s:file:%s", projectId, fileId)
}

// keysFromFileRequests extracts the file IDs from the given file requests
// and builds the Redis keys for them
func keysFromFileRequests(projectId string, filesRequest domain.FilesRequest, outKeys *[]string) {
	for _, fileId := range filesRequest.GetFileIDs() {
		*outKeys = append(*outKeys, buildRedisKeyForFile(projectId, fileId))
	}
}
