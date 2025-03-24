package module1

import "fmt"

func reverseString() {
	person := []byte{'h', 'e', 'e', 'l', 'o'}

	for i, j := 0, len(person)-1; i < j; i, j = i+1, j-1 {
		person[i], person[j] = person[j], person[i]
	}

	fmt.Printf("%c", person)

}
