package module1

import (
	"fmt"
	"slices"
)

func findEvenNumbers(digits []int) []int {
	hashmap := make(map[int]int)
	for _, value := range digits {
		hashmap[value]++
	}
	result := []int{}
	for i := 100; i < 999; i++ {
		num := i
		nummap := make(map[int]int)
		if num%2 == 0 {
			count := 0
			counter := 0
			for num > 0 {
				temp := num % 10
				num = num / 10
				if !slices.Contains(digits, temp) {
					continue
				} else {
					counter++
					nummap[temp]++
				}
			}
			fmt.Println("num", nummap, i)
			fmt.Println("hashmap", hashmap)
			if counter == 3 {
				for key, value := range nummap {
					fmt.Println("key", key, value)
					if hashmap[key] >= value {
						count++
						fmt.Println("", count)
					}
				}
				if count == len(nummap) {
					result = append(result, i)
				}
			}
		}
	}
	return result
}
