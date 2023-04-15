package util

import "strings"

func NotBlank(str string) bool {
	trim := strings.Trim(str, " ")
	return len(trim) > 0
}

func NotBlankThen(str string, then func(str string)) {
	if NotBlank(str) {
		then(str)
	}
}
