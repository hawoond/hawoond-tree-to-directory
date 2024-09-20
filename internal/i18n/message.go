package i18n

var Messages = map[string]map[string]string{
	"en": {
		"usage":            "Usage: generate-structure -input=FILE_PATH [-indent=INDENT_SIZE] [-language=LANGUAGE] [--add-package]",
		"input_desc":       "Specify the path to the structure file (-input, -i)",
		"indent_desc":      "Specify the indent size (default: 4) (-indent, -n)",
		"language_desc":    "Specify the programming language (e.g., go, python) (-language, -lang, -l)",
		"add_package_desc": "Add package declarations to Go files (use with -language=go) (-add-package, -p)",
		"open_error":       "Cannot open the structure file: %v",
		"dir_create_err":   "Error creating directory: %v",
		"file_create_err":  "Error creating file: %v",
		"read_error":       "Error reading structure file: %v",
	},
	"ko": {
		"usage":            "사용법: generate-structure -input=파일경로 [-indent=들여쓰기 칸 수] [-language=언어] [--add-package]",
		"input_desc":       "구조 파일의 경로를 지정합니다 (-input, -i)",
		"indent_desc":      "들여쓰기 칸 수를 지정합니다 (기본값: 4) (-indent, -n)",
		"language_desc":    "프로그래밍 언어를 지정합니다 (예: go, python) (-language, -lang, -l)",
		"add_package_desc": "Go 파일에 패키지 선언을 추가합니다 (-language=go와 함께 사용) (-add-package, -p)",
		"open_error":       "구조 파일을 열 수 없습니다: %v",
		"dir_create_err":   "디렉토리를 생성하는 중 오류 발생: %v",
		"file_create_err":  "파일을 생성하는 중 오류 발생: %v",
		"read_error":       "구조 파일을 읽는 중 오류 발생: %v",
	},
}
