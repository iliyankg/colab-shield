package common

import (
	"fmt"

	"github.com/iliyankg/colab-shield/backend/common/request"
)

// buildRedisKeyForFile builds the Redis key for the given project and file IDs
func buildRedisKeyForFile(projectId string, fileId string) string {
	return fmt.Sprintf("project:%s:file:%s", projectId, fileId)
}

// keysFromFileRequests extracts the file IDs from the given file requests
// and builds the Redis keys for them
func keysFromFileRequests(projectId string, filesRequest request.FilesRequest, outKeys *[]string) {
	for _, fileId := range filesRequest.GetFilesIds() {
		*outKeys = append(*outKeys, buildRedisKeyForFile(projectId, fileId))
	}
}
