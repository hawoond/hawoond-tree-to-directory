package main

import (
	"fmt"
	"os"

	"github.com/hawoond/hawoond-tree-to-directory/internal/flags"
	"github.com/hawoond/hawoond-tree-to-directory/internal/generator"
	"github.com/hawoond/hawoond-tree-to-directory/internal/i18n"
	"github.com/hawoond/hawoond-tree-to-directory/internal/util"
)

func main() {
	locale := util.DetectLocale()
	if _, exists := i18n.Messages[locale]; !exists {
		locale = "en"
	}

	messages := i18n.Messages[locale]

	config := flags.ParseFlags(messages)

	err := generator.GenerateStructure(config.InputFile, config.IndentSize, config.Language, config.AddPackage, messages)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
