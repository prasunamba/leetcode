package module1

import (
	"fmt"
	"strings"
)

func First_Non_Repeating_Character() {
	name := "yellow"
	name = strings.ReplaceAll(name, " ", "")
	runes := []rune(name)
	map1 := make(map[rune]int)
	var First_Non_Repeating_Character rune
	for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
		map1[runes[i]]++
		if i != j {
			map1[runes[j]]++
		}

	}
	set := false
	for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
		if map1[runes[i]] == 1 {
			First_Non_Repeating_Character = runes[i]
			set = true
			if !set {
				if i != j && map1[runes[j]] == 1 { // Check `j` only if different from `i`
					First_Non_Repeating_Character = runes[j]
					break
				}
			}
			break
		}

	}
	fmt.Printf(" main %c \n", First_Non_Repeating_Character)
}
