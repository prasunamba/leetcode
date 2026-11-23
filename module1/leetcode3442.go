package module1

import "fmt"

func MaxDifference(s string) int {
	smap := make(map[rune]int)
	for _, value := range s {
		smap[value]++
	}
	evenfreq := len(s)
	oddfreq := 0
	for _, value := range smap {
		if value < evenfreq && value%2 == 0 {
			evenfreq = value
		} else if value > oddfreq && value%2 != 0 {
			oddfreq = value
		}

	}
	fmt.Println("", oddfreq, evenfreq)
	return oddfreq - evenfreq
}
