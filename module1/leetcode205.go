package module1

import "fmt"

func isIsomorphic(s string, t string) bool {
	srune := []rune(s)
	trune := []rune(t)
	hashmap := make(map[rune]rune)
	for idx, value := range srune { // e g g
		if _, exists := hashmap[value]; !exists {
			fmt.Println("if", exists, value, hashmap[value])
			hashmap[value] = trune[idx]
			fmt.Printf("%c", hashmap)
		} else if hashmap[value] != trune[idx] {
			fmt.Printf("%c,%c", value, hashmap[value])
			return false
		}
	}
	fmt.Println("")
	for i, j := 0, len(trune)-1; i <= j; i, j = i+1, j-1 {
		trune[i], trune[j] = srune[i], srune[j]
		fmt.Printf("%c", trune)
	}
	return string(trune) == string(srune)
}
func IsIsomorphic_optimized(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]rune)
	tMap := make(map[rune]rune)

	for i, sChar := range s {
		tChar := rune(t[i])
		fmt.Printf("%c%c\n ", sChar, tChar)
		// Check if s → t mapping is consistent
		if mapped, exists := sMap[sChar]; exists {
			if mapped != tChar {
				return false
			}
		} else {
			sMap[sChar] = tChar
		}

		// Check if t → s mapping is consistent
		if mapped, exists := tMap[tChar]; exists {
			if mapped != sChar {
				return false
			}
		} else {
			tMap[tChar] = sChar
		}
	}

	return true
}
