package util

import "os"

func DetectLocale() string {
	locale := os.Getenv("LC_ALL")
	if locale == "" {
		locale = os.Getenv("LC_MESSAGES")
	}
	if locale == "" {
		locale = os.Getenv("LANG")
	}
	if locale == "" {
		locale = "en"
	} else {
		if len(locale) >= 2 {
			locale = locale[:2]
		}
	}
	return locale
}
