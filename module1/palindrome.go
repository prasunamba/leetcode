package module1

import "fmt"

func palindrome() {
	name := "abckba"
	runes := []rune(name)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		if runes[i] != runes[j] {
			fmt.Println("not a palindrome")
		}

	}
}
