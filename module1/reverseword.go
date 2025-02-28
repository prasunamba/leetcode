package module1

import "fmt"

func ReverseString(name string) {
	nam := []rune(name) // convert string to slice of runes
	for i := 0; i < len(nam)/2; i++ {
		nam[i], nam[len(nam)-i-1] = nam[len(name)-i-1], nam[i]
	}
	fmt.Println(string(nam))
}

// 0 1 2
// D G P

// i:=0  D,P=P,D      PGD
// i=1   G,G=G,G      PGD
// i=2   P,
