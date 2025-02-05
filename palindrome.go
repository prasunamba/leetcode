package main

import "fmt"

func palindrome() {
	name := "abckba"
	for i, value := range name {
		fmt.Println(value, rune(name[len(name)-i-1]))
		if value != rune(name[len(name)-i-1]) {
			fmt.Println("not a palindrome ")
			return
		}
	}
	fmt.Println("a palindrome ")
}
