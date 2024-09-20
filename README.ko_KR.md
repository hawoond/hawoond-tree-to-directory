# hawwond Tree Generator

텍스트 파일로 정의된 디렉토리 구조를 실제 폴더와 파일로 생성합니다

## 설치 방법

Go 1.16 이상이 필요합니다.

### 최신 버전 설치

```bash
go install github.com/hawoond/hawoond-tree-generator/cmd/htg@latest
```

## 사용 방법
 
Directory Tree를 파일 형태로 저장 후 argument로 파일 경로를 입력하여 실행

### 기본 사용
```bash
htg -input=tree.txt
```

### 언어 지정
- go
```
htg -input=tree.txt -language=go
```
- package 선언 추가
```
htg -input=tree.txt -language=go -add-package
```

### 파일 포맷 예시

```
project/
├── cmd/
│   └── app/
│       └── main.go
└── README.md
```
