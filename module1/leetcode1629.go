package module1

import (
	"fmt"
	"sort"
)

func SlowestKey(releaseTimes []int, keysPressed string) byte {
	keymap := make(map[int]int)
	for index, value := range keysPressed {
		if index == 0 {
			keymap[int(value)] = releaseTimes[index]
		} else {
			num := releaseTimes[index] - releaseTimes[index-1]
			if keymap[int(value)] < num {
				keymap[int(value)] = num
			}
		}
	}
	fmt.Println(keymap)
	max := 0
	result := []int{}
	for _, value := range keymap {
		if value > max {
			max = value
		}
	}
	for key, value := range keymap {
		if value == max {
			result = append(result, key)
		}
	}
	sort.Ints(result)
	fmt.Println(result)
	return byte(result[len(result)-1])
}
