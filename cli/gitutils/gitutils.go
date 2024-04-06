package gitutils

import (
	"os/exec"
	"strings"

	"github.com/rs/zerolog"
	
)

// GetGitBlobHashes returns the git hashes for the files passed as arguments.
// Expects the files to be in the format "path/to/file" from the root of the repository.
func GetGitBlobHashes(log *zerolog.Logger, filesToProcess []string) ([]string, error) {
	args := []string{"hash-object"}
	args = append(args, filesToProcess...)

	executedCommand := exec.Command("git", args...)
	output, err := executedCommand.Output()
	if err != nil {
		log.Err(err).Msg(string(output))
		return nil, err
	}
	stringifiedOutput := string(output)
	hashes := strings.Split(stringifiedOutput, "\n")[0:len(filesToProcess)] // Remove the last empty string

	return hashes, nil
}

// GetGitBlobHEADHashes returns the git HEAD hashes for the files passed as arguments.
// Expects the files to be in the format "path/to/file" from the root of the repository.
func GetGitBlobHEADHashes(log *zerolog.Logger, filesToProcess []string) ([]string, error) {
	args := []string{"rev-parse"}
	for _, file := range filesToProcess {
		args = append(args, "HEAD:"+file)
	}

	executedCommand := exec.Command("git", args...)
	output, err := executedCommand.Output()
	if err != nil {
		log.Err(err).Msg(string(output))
		return nil, err
	}
	stringifiedOutput := string(output)
	hashes := strings.Split(stringifiedOutput, "\n")[0:len(filesToProcess)] // Remove the last empty string

	return hashes, nil
}
