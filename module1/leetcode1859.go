package module1

import (
	"fmt"
	"strings"
)

func SortSentence(s string) string {
	words := strings.Fields(s)
	result := make([]string, len(words))
	fmt.Println("words", words)
	for _, value := range words {
		// l := len(value)
		fmt.Println("", value, len(value), value[len(value)-1]-'0', string(value[:len(value)-1]))
		result[value[len(value)-1]-'0'-1] = string(value[:len(value)-1])

	}
	fmt.Println("", result)

	return strings.Join(result, " ") // dont use the loop to add words
}
