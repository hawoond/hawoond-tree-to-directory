package flags

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	InputPath  string
	OutputPath string
	IndentSize int
	Language   string
	AddPackage bool
	Reverse    bool
	Include    []string
	Exclude    []string
}

func ParseFlags(messages map[string]string) *Config {
	var config Config
	var includePatterns string
	var excludePatterns string

	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flagSet.StringVar(&config.InputPath, "input", "", messages["input_desc"])
	flagSet.StringVar(&config.OutputPath, "output", "", messages["output_desc"])
	flagSet.IntVar(&config.IndentSize, "indent", 4, messages["indent_desc"])
	flagSet.StringVar(&config.Language, "language", "", messages["language_desc"])
	flagSet.BoolVar(&config.AddPackage, "add-package", false, messages["add_package_desc"])
	flagSet.BoolVar(&config.Reverse, "reverse", false, messages["reverse_desc"])
	flagSet.StringVar(&includePatterns, "include", "", messages["include_desc"])
	flagSet.StringVar(&excludePatterns, "exclude", "", messages["exclude_desc"])

	flagSet.StringVar(&config.InputPath, "i", "", messages["input_desc"])
	flagSet.StringVar(&config.OutputPath, "o", "", messages["output_desc"])
	flagSet.IntVar(&config.IndentSize, "n", 4, messages["indent_desc"])
	flagSet.StringVar(&config.Language, "l", "", messages["language_desc"])
	flagSet.StringVar(&config.Language, "lang", "", messages["language_desc"])
	flagSet.BoolVar(&config.AddPackage, "p", false, messages["add_package_desc"])
	flagSet.BoolVar(&config.Reverse, "r", false, messages["reverse_desc"])
	flagSet.StringVar(&includePatterns, "I", "", messages["include_desc"])
	flagSet.StringVar(&excludePatterns, "E", "", messages["exclude_desc"])

	flagSet.Usage = func() {
		fmt.Println(messages["usage"])
		flagSet.PrintDefaults()
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			flagSet.Usage()
			os.Exit(0)
		}
	}

	flagSet.Parse(os.Args[1:])

	if config.Reverse {
		if config.InputPath == "" || config.OutputPath == "" {
			fmt.Println(messages["reverse_usage"])
			os.Exit(1)
		}
	} else {
		if config.InputPath == "" {
			fmt.Println(messages["usage"])
			os.Exit(1)
		}
	}

	if includePatterns != "" {
		config.Include = strings.Split(includePatterns, ",")
	}
	if excludePatterns != "" {
		config.Exclude = strings.Split(excludePatterns, ",")
	}

	return &config
}
