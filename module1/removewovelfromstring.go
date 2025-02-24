package module1

import "fmt"

func RemoveVowel() {
	vowel := "aeiouAEIOU"
	vowels := []rune(vowel)
	senten := " i am prasunamba"
	result := ""
	for _, c := range senten {
		isvowel := false
		for _, v := range vowels {
			if c == v {
				isvowel = true
				break
			}
		}
		if !isvowel {
			result += string(c)
		}
	}
	fmt.Println("", result)
}
