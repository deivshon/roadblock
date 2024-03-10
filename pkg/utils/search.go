package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func SearchFilesRecursive(dirPath string, needleFilename string) ([]string, error) {
	pathData, err := os.Stat(dirPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []string{}, nil
		}

		return nil, fmt.Errorf("could not stat `%v`: %w", dirPath, err)
	}
	if !pathData.IsDir() {
		return nil, fmt.Errorf("recursive search was passed a non-directory: %v", dirPath)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("could not list files in dir: %v", dirPath)
	}

	results := make([]string, 0)
	for _, e := range entries {
		completeEntryPath := filepath.Join(dirPath, e.Name())
		if !e.IsDir() && e.Name() == needleFilename {
			results = append(results, completeEntryPath)
		} else if e.IsDir() {
			recursiveResults, err := SearchFilesRecursive(completeEntryPath, needleFilename)
			if err != nil {
				return nil, err
			}

			results = append(results, recursiveResults...)
		}
	}

	return results, nil
}
