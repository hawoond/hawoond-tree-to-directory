package generator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hawoond/hawoond-tree-to-directory/internal/util"
)

func GenerateStructure(inputFile string, indentSize int, language string, addPackage bool, includePatterns, excludePatterns []string, messages map[string]string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf(messages["open_error"], err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pathStack []string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		leadingChars := util.CountLeadingChars(line)
		depthLevel := leadingChars / indentSize

		name := strings.TrimSpace(util.RemovePrefixes(line))

		if depthLevel < len(pathStack) {
			pathStack = pathStack[:depthLevel]
		}

		pathStack = append(pathStack, name)

		fullPath := filepath.Join(pathStack...)

		isDir := strings.HasSuffix(name, "/") || strings.HasSuffix(name, "\\")

		if !util.ShouldInclude(fullPath, includePatterns, excludePatterns, isDir) {
			continue
		}

		if isDir {
			err = util.CreateDirectory(fullPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf(messages["dir_create_err"], err)
			}
		} else {
			err = util.CreateFileWithContent(fullPath, language, addPackage)
			if err != nil {
				return fmt.Errorf(messages["file_create_err"], err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf(messages["read_error"], err)
	}

	return nil
}

func SaveStructureToFile(rootDir string, outputFile string, indentSize int, includePatterns, excludePatterns []string, messages map[string]string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf(messages["write_error"], err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	rootDirAbs, err := filepath.Abs(rootDir)
	if err != nil {
		return err
	}

	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		pathAbs, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		if pathAbs == rootDirAbs {
			return nil
		}

		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}

		isDir := info.IsDir()
		if !util.ShouldInclude(relPath, includePatterns, excludePatterns, isDir) {
			if isDir {
				return filepath.SkipDir
			}
			return nil
		}

		depth := len(strings.Split(relPath, string(os.PathSeparator))) - 1

		indent := strings.Repeat(strings.Repeat(" ", indentSize), depth)

		name := filepath.Base(path)
		if isDir {
			name += "/"
		}

		line := fmt.Sprintf("%s%s\n", indent, name)
		_, err = writer.WriteString(line)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf(messages["write_error"], err)
	}

	return nil
}
