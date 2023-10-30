package file

import (
	"embed"
	"fmt"
	"io"
)

func ReadContent(fs embed.FS, filePath string) ([]byte, error) {
	sourceFile, err := fs.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", filePath, err)
	}

	fileContent, err := io.ReadAll(sourceFile)
	if err != nil {
		return nil, fmt.Errorf("could not read %s file content: %v", filePath, err)
	}

	return fileContent, nil
}
