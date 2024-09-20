# hawoond Tree to Directory

Generate actual folders and files from a directory structure defined in a text file.

## Purpose

This program reads a directory tree structure expressed in text and creates directories and files in the actual file system accordingly. Useful when feeling drowsy.

## Installation

Requires Go 1.16 or higher.

### Installation via `go install`

```bash
go install github.com/hawoond/hawoond-tree-to-directory/cmd/htd@latest
```

After installation, the executable file will be located in `$GOPATH/bin` or `$GOBIN`.

## Features and Usage

### Basic Usage

```bash
htd -input=tree.txt
```

Shortcut:

```bash
htd -i=tree.txt
```

`tree.txt` is a text file that describes the directory and file structure you want to create.

### Command-Line Options

- `-input`, `-i`: Specify the path to the structure file or directory (required).
- `-output`, `-o`: Specify the path to the output file (required in `-reverse` mode).
- `-indent`, `-n`: Specify the number of indent spaces (including whitespace) (default: 4).
- `-language`, `-lang`, `-l`: Specify the programming language (e.g., `go`, `python`).
- `-add-package`, `-p`: For Go language, adds package declarations to `.go` files (use with `-language=go`).
- `-reverse`, `-r`: Reads the directory structure and saves it as text.
- `-include`, `-I`: Specify patterns of files/directories to include, separated by commas.
- `-exclude`, `-E`: Specify patterns of files/directories to exclude, separated by commas.

### Detailed Features

1. **Directory Structure Generation**

   - Creates directories and files in the actual file system based on the structure defined in a text file.
   - Supports various indentation styles by adjusting the number of indent spaces.

2. **Directory Structure Saving**

   - Saves the existing directory structure as a text file.
   - Represents the tree structure by specifying the number of indent spaces.

3. **Programming Language Support**

   - Allows specifying a programming language using the `-language` option.
   - Currently supports special handling for the `go` language.

4. **Adding Go Package Declarations**

   - When used with `-language=go` and `-add-package`, adds package declarations to `.go` files.
   - The package name is generated based on the name of the directory containing the file.

5. **Locale Detection and Multilingual Support**

   - Detects the system locale and outputs messages in the corresponding language.
   - Supported languages: English (`en`), Korean (`ko`).

6. **Shortcut Flags Support**

   - Supports shortcut flags for command-line options.
   - Examples: `-input` ➔ `-i`, `-language` ➔ `-l`

7. **File and Directory Filtering**

   - Use `-include` or `-I` options to specify patterns of files/directories to include.
   - Use `-exclude` or `-E` options to specify patterns of files/directories to exclude.
   - Patterns use Glob patterns and are separated by commas.

### Usage Examples

#### Directory Generation

##### Structure File (`tree.txt`)

```
project/
├── cmd/
│   └── app/
│       └── main.go
└── README.md
```

##### Running the Program

```bash
htd -i=tree.txt -l=go -p
```

##### Result

- `project/cmd/app/main.go` file:

  ```go
  package app

  ```

- `project/pkg/utils/helper.go` file:

  ```go
  package utils

  ```

- `project/README.md` file created (empty content).

#### Saving Directory Structure

##### Running the Program

```bash
htd -r -i=project -o=tree.txt -n=4
```

##### Result

- The structure of the `project` directory is saved to `tree.txt`.

#### Excluding Specific Files During Directory Generation

```bash
htd -i=tree.txt -E="*.md,*.txt"
```

- Generates the directory excluding `.md` and `.txt` files.

#### Including Specific Directories When Saving Structure

```bash
htd -r -i=project -o=tree.txt -I="cmd,pkg"
```

- Saves the structure including only the `cmd` and `pkg` directories.

### Viewing Help

```bash
htd -h
```

or

```bash
htd --help
```

## Notes

- **Indentation Settings**

  - The default number of indent spaces is 4. You can change it using the `-indent` option. (Default: 4)
  - If using tab characters for indentation, set `-indent=1`.

- **Locale Settings**

  - The program detects the locale through the `LANG`, `LC_ALL`, and `LC_MESSAGES` environment variables.
  - If the locale is not supported, the default is English (`en`).

- **Pattern Matching**

  - The patterns for the `-include` and `-exclude` options apply to the **base name** of files or directories.
  - Patterns use [Glob patterns](https://golang.org/pkg/path/filepath/#Match).
  - Example: `*.go` matches all files with the `.go` extension.

- **Package Name Rules**

  - In Go, package names should be in lowercase and can only contain letters, numbers, and underscores (`_`).
  - When converting directory names to package names, special characters are removed or replaced with underscores as needed.

## License

This project is licensed under the **MIT License**. For more details, see the `LICENSE` file.

---
