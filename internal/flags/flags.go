package flags

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	InputFile  string
	OutputFile string
	IndentSize int
	Language   string
	AddPackage bool
	Reverse    bool
}

func ParseFlags(messages map[string]string) *Config {
	var config Config

	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flagSet.StringVar(&config.InputFile, "input", "", messages["input_desc"])
	flagSet.StringVar(&config.OutputFile, "output", "", messages["output_desc"])
	flagSet.IntVar(&config.IndentSize, "indent", 4, messages["indent_desc"])
	flagSet.StringVar(&config.Language, "language", "", messages["language_desc"])
	flagSet.BoolVar(&config.AddPackage, "add-package", false, messages["add_package_desc"])
	flagSet.BoolVar(&config.Reverse, "reverse", false, messages["reverse_desc"])

	flagSet.StringVar(&config.InputFile, "i", "", messages["input_desc"])
	flagSet.StringVar(&config.OutputFile, "o", "", messages["output_desc"])
	flagSet.IntVar(&config.IndentSize, "n", 4, messages["indent_desc"])
	flagSet.StringVar(&config.Language, "l", "", messages["language_desc"])
	flagSet.StringVar(&config.Language, "lang", "", messages["language_desc"])
	flagSet.BoolVar(&config.AddPackage, "p", false, messages["add_package_desc"])
	flagSet.BoolVar(&config.Reverse, "r", false, messages["reverse_desc"])

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
		if config.InputFile == "" {
			fmt.Println(messages["reverse_usage"])
			os.Exit(1)
		}
		if config.OutputFile == "" {
			fmt.Println(messages["reverse_usage"])
			os.Exit(1)
		}
	} else {
		if config.InputFile == "" {
			fmt.Println(messages["usage"])
			os.Exit(1)
		}
	}

	return &config
}
