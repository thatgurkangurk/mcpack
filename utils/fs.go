package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func CheckOrCreateFolder(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		return nil
	}

	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	if len(files) > 0 {
		return errors.New("directory is not empty")
	}

	return nil
}

func EmptyDirectory(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())

		if file.IsDir() {
			err := EmptyDirectory(filePath)
			if err != nil {
				return err
			}
			err = os.Remove(filePath)
			if err != nil {
				return err
			}
		} else {
			err := os.Remove(filePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
