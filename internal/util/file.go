package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func ShouldInclude(path string, includePatterns, excludePatterns []string, isDir bool) bool {
	if isDir && !strings.HasSuffix(path, string(os.PathSeparator)) {
		path += string(os.PathSeparator)
	}

	for _, pattern := range excludePatterns {
		matched, err := filepath.Match(pattern, filepath.Base(path))
		if err == nil && matched {
			return false
		}
	}

	if len(includePatterns) == 0 {
		return true
	}

	for _, pattern := range includePatterns {
		matched, err := filepath.Match(pattern, filepath.Base(path))
		if err == nil && matched {
			return true
		}
	}

	return false
}
