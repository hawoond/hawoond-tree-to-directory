package generator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hawoond/hawoond-tree-generator/internal/util"
)

func GenerateStructure(inputFile string, indentSize int, language string, addPackage bool, messages map[string]string) error {

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
