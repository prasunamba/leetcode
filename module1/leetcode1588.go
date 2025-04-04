package module1

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			// fmt.Println("arr", arr[i:j+1])
			if len(arr[i:j+1])%2 != 0 {

				temp := 0
				tempslice := arr[i : j+1]
				for k := 0; k < len(arr[i:j+1]); k++ {
					fmt.Println("tempslice[k]", tempslice[k])
					temp += tempslice[k]
				}
				fmt.Println("temp", tempslice, temp)
				sum += temp
				// fmt.Println("sum", sum)
			}
		}
	}
	return sum
}
