package module1

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	wordslice := strings.Fields(s)
	fmt.Println("", wordslice)
	result := []string{}
	for j := len(wordslice) - 1; j >= 0; j-- {
		result = append(result, wordslice[j])
	}
	fmt.Println("", result)
	return strings.Join(result, " ")
}
