package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindFile(directory, filename string) (string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Printf("unable to find the file %q", filename)
		return "", err
	}

	for _, file := range files {
		// Strip the extension from the file
		fileWithoutExtension := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		if fileWithoutExtension == filename {
			// Return the full file path if match is found
			sanatisedFilePath := strings.Replace(filepath.Join(directory, file.Name()), `\`, "/", -1)
			return sanatisedFilePath, nil
		}
	}
	return "", nil
}
