package module1

func BubbleSort(list []int) []int {
	// time complexity -o(n2)
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ { // at the end of ith iteration  the ith  largest element will go to last of the array
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func OptimizedBsort(slice []int) []int {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}
		}
		if !swapped { // If no swaps in this pass, stop early
			break
		}
	}
	return slice
}

/* func main() {
	var num int
	var choice string
	list := []int{}
	for {
		fmt.Println("Enter a number : ")
		fmt.Scan(&num)
		list = append(list, num)
		fmt.Println(list)
		fmt.Println(module1.BubbleSort(list))
		fmt.Println("DO you want to continue(y/n)  : ")
		fmt.Scan(&choice)
		if strings.ToLower(choice) == "n" {
			fmt.Println("exiting the program ")
			break
		}

	}

} */
