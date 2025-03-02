package module1

import (
	"strings"
	"unicode"
)

func isPalindrome(name string) bool {
	if name == "" {
		return true
	}
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ToLower(name)
	var newname []rune
	for _, char := range name {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			newname = append(newname, char)
		}
	}
	for i, j := 0, len(newname)-1; i < j; i, j = i+1, j-1 {
		if newname[i] != newname[j] {

			return false
		}
	}
	return true
}
