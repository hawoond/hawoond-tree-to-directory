package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDirectory(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func CreateFileWithContent(path, language string, addPackage bool) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if language == "go" && addPackage && filepath.Ext(path) == ".go" {
		packageName := filepath.Base(dir)
		content := fmt.Sprintf("package %s\n\n", packageName)
		_, err = f.WriteString(content)
		if err != nil {
			return err
		}
	}

	return nil
}
