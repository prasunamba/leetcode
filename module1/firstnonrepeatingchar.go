package module1

import (
	"fmt"
	"strings"
)

func First_Non_Repeating_Character(name string) rune {
	name = strings.ReplaceAll(name, " ", "")
	runes := []rune(name)

	// Step 1: Count frequencies
	freq := make(map[rune]int)
	for _, ch := range runes {
		freq[ch]++
	}

	// Step 2: Find first character with frequency 1
	for _, ch := range runes {
		if freq[ch] == 1 {
			fmt.Printf("First non-repeating character: %c\n", ch)
			return ch
		}
	}

	fmt.Println("No non-repeating character found")
	return ' '
}
