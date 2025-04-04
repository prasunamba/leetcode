package module1

import "fmt"

func Pattern1() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}
func Pattern2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(i+1, " ")
		}
		fmt.Println("")
	}
}
func Pattern3() {
	for i := 5; i > 0; i-- {
		for j := 0; j < 5; j++ {
			fmt.Print(j+1, " ")
		}
		fmt.Println("")
	}
}
func Pattern4() {
	for i := 5; i > 0; i-- {
		for j := 0; j < 5; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println("")
	}
}
func Pattern5() {
	for i := 5; i > 0; i-- {
		for j := 5; j > 0; j-- {
			fmt.Print(j, " ")
		}
		fmt.Println("")
	}
}
func Pattern6() {
	count := 1
	for i := 0; i < 5; i++ {
		for j := 1; j < 6; j++ {
			fmt.Print(count, " ")
			count++
		}
		fmt.Println("")
	}
}
func Pattern7() {
	count := 1
	for i := 0; i < 5; i++ {
		for j := 1; j < 6; j++ {
			fmt.Print(count, " ")
			count += 2
		}
		fmt.Println("")
	}
}
func Pattern8() {
	count := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			count += 2
			fmt.Print(count, " ")
		}
		fmt.Println("")
	}
}
func Pattern9() {
	for i := 1; i < 6; i++ {
		for j := 1; j < 6; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println("")
	}
}
func Pattern10() {
	for i := 1; i < 6; i++ {
		for j := 1; j < 4; j++ {
			fmt.Print(j, i, " ")
		}
		fmt.Println("")
	}
}
func Pattern11() {
	for i := 1; i < 6; i++ {
		for j := 1; j < 4; j++ {
			fmt.Print(i, j, " ")
		}
		fmt.Println("")
	}
}
func Pattern12() {
	for i := 1; i < 6; i++ {
		count := i
		for j := 1; j < 6; j++ {
			fmt.Print(count, " ")
			count += 5
		}
		fmt.Println("")
	}
}
func Pattern13() {
	for i := 1; i < 6; i++ {
		x := i
		y := 5 - i + 1
		for j := 1; j < 6; j++ {
			if j%2 == 1 {
				fmt.Print(x, " ") //odd
			} else {
				fmt.Print(y, " ")
			}
			x += 5
			y += 5
		}
		fmt.Println("")
	}
}
func Pattern14() {
	for i := 0; i < 5; i++ {

		for j := 1; j < 6; j++ {
			fmt.Print(j*5-i, " ")
		}
		fmt.Println("")
	}
}
