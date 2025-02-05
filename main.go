package main

// func main() {
// 	// fmt.Println("Hello World")
// 	// Helloworld()
// 	fmt.Println(BubbleSort())
// 	f := 3.14
// 	fmt.Println("f", f)

// }
func BubbleSort() []int {
	list := []int{12, 32, 1, 45, -1, -23}
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}
