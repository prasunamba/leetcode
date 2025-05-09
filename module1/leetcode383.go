package module1

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	mapransom := make(map[rune]int)
	mapmagazine := make(map[rune]int)
	for _, value := range ransomNote {
		mapransom[value]++
	}
	for _, value := range magazine {
		mapmagazine[value]++
	}
	for _, value := range ransomNote {
		fmt.Println("", mapmagazine[value], mapransom[value], value)
		if mapmagazine[value] >= mapransom[value] {
			continue
		} else {
			return false
		}
	}
	return true
}
