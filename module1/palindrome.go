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
func optimizedpalindrome(name string) bool {
	for i, j := 0, len(name)-1; i <= j; i, j = i+1, j-1 {
		if name[i] != name[j] {
			return false
		}
	}
	return true
}
