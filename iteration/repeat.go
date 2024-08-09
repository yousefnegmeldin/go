package iteration

import "strings"

func Repeat(str string) (result string) {
	result = strings.Repeat(str, 5)
	return result
}
