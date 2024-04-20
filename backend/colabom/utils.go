package colabom

import "fmt"

// buildRedisKeyForFile builds the Redis key for the given project and file IDs
func BuildRedisKeyForFile(projectId string, fileId string) string {
	return fmt.Sprintf("project:%s:file:%s", projectId, fileId)
}
