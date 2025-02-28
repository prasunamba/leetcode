package module1

import (
	"fmt"
	"strings"
)

func Reversestring() {
	name := "hello i am ok"
	words := strings.Fields(name)
	for i, word := range words {
		words[i] = reverse(word)
	}
	result := strings.Join(words, " ")
	fmt.Println("", result)

}
func reverse(name string) string {
	runeslice := []rune(name)
	// namespace := strings.Fields(name)
	for i, j := 0, len(runeslice)-1; i < j; i, j = i+1, j-1 {
		runeslice[i], runeslice[j] = runeslice[j], runeslice[i]
	}
	return string(runeslice)
}
