package utils

func BuildFileToHashMap(files []string, hashes []string) map[string]string {
	if len(files) != len(hashes) {
		return nil
	}

	fileToHash := make(map[string]string)
	for i, file := range files {
		fileToHash[file] = hashes[i]
	}
	return fileToHash
}
