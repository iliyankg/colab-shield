package gitutils

import (
	"os/exec"
	"strings"
)

// GetGitBlobHashes returns the git hashes for the files passed as arguments.
func GetGitBlobHashes(filesToProcess []string) ([]string, error) {
	args := []string{"hash-object"}
	args = append(args, filesToProcess...)

	executedCommand := exec.Command("git", args...)
	output, err := executedCommand.Output()
	if err != nil {
		return nil, err
	}
	stringifiedOutput := string(output)
	hashes := strings.Split(stringifiedOutput, "\n")[0:len(filesToProcess)] // Remove the last empty string

	return hashes, nil
}
