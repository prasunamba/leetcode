package module1

import "fmt"

func IndexofSubstring() {
	bigstring := "helloc"
	find := "c"
	index := -1

	for i := 0; i <= len(bigstring)-len(find); i++ {
		fmt.Println("", i, bigstring[i:i+len(find)])
		if bigstring[i:i+len(find)] == find {
			index = i
			fmt.Println("-", index)
		}
	}
}
