package module1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LongestWord() {
	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	name := strings.TrimSpace(reader)
	list := strings.Fields(name)
	maxLen := 0
	var longestWords []string
	for _, word := range list {
		if len(word) > maxLen {
			maxLen = len(word)
			longestWords = []string{word} // Reset with the new longest word
		} else if len(word) == maxLen {
			longestWords = append(longestWords, word) // Add if same length
		}
	}
	fmt.Println("long", longestWords)
}
