# hawoond Tree Generator

Generate actual folders and files from a directory structure defined in a text file.

## Purpose

This program reads a directory tree structure expressed in text and creates directories and files in the actual file system accordingly. It is useful when you need to quickly set up a directory structure, especially when feeling drowsy.

## Installation

Requires Go 1.16 or higher.

### Installation via `go install`

```bash
go install github.com/hawoond/hawoond-tree-generator/cmd/htg@latest
```

After installation, the executable file will be located in `$GOPATH/bin` or `$GOBIN`.

## Features and Usage

### Basic Usage

```bash
htg -input=tree.txt
```

Shortcut:

```bash
htg -i=tree.txt
```

`tree.txt` is a text file that describes the directory and file structure you want to create.

### Command-Line Options

- `-input`, `-i`: Specifies the path to the structure file (required).
- `-indent`, `-n`: Specifies the number of indent spaces (including whitespace) (default: 4).
- `-language`, `-lang`, `-l`: Specifies the programming language (e.g., `go`, `python`).
- `-add-package`, `-p`: For Go language, adds package declarations to `.go` files (use with `-language=go`).

### Detailed Features

1. **Directory Structure Generation**

   - Creates directories and files in the actual file system based on the structure defined in a text file.
   - Supports various indentation styles by adjusting the number of indent spaces.

2. **Programming Language Support**

   - Allows specifying a programming language using the `-language` option.
   - Currently supports special handling for the `go` language.

3. **Adding Go Package Declarations**

   - When used with `-language=go` and `-add-package`, adds package declarations to `.go` files.
   - The package name is based on the name of the directory containing the file.

4. **Locale Detection and Multilingual Support**

   - Detects the system locale and outputs messages in the corresponding language.
   - Supported languages: English (`en`), Korean (`ko`).

### Usage Example

#### Structure File (`tree.txt`)

```
project/
├── cmd/
│   └── app/
│       └── main.go
├── pkg/
│   └── utils/
│       └── helper.go
└── README.md
```

#### Running the Program

```bash
htg -i=tree.txt -l=go -p
```

#### Result

- `project/cmd/app/main.go` file:

  ```go
  package app

  ```

- `project/pkg/utils/helper.go` file:

  ```go
  package utils

  ```

- `project/README.md` file created (empty content).

### Viewing Help

```bash
htg -h
```

or

```bash
htg --help
```

## Notes

- **Indentation Settings**

  - The default number of indent spaces is 4. You can change it using the `-indent` option.
  - If using tab characters for indentation, set `-indent=1`.

- **Locale Settings**

  - The program detects the locale through the `LANG`, `LC_ALL`, and `LC_MESSAGES` environment variables.
  - If the locale is not supported, the default is English (`en`).

- **Package Name Rules**

  - In Go, package names should be in lowercase and can only contain letters, numbers, and underscores (`_`).
  - When converting directory names to package names, special characters are removed or replaced with underscores as needed.

## License

This project is licensed under the **Apache License 2.0**. For more details, see the `LICENSE` file.