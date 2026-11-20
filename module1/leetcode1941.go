package module1

import "fmt"

func AreOccurrencesEqual(s string) bool {
	strmap := make(map[rune]int)
	for _, value := range s {
		strmap[value]++
	}
	values := []int{}
	for _, value := range strmap {
		values = append(values, value)
	}
	fmt.Println("", values)
	for i := 0; i < len(values)-1; i++ {
		if values[i] != values[i+1] {
			return false
		}
	}
	return true
}
