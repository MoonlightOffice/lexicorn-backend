package util

import "strings"

func Trim(text string) string {
	original := text

	cutsets := []string{`"`, ` `, `　`, `	`}

	for _, cutset := range cutsets {
		text = strings.Trim(text, cutset)
	}

	if text != original {
		text = Trim(text)
	}

	return text
}
