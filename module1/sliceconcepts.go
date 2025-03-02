package module1

func SlicesConcepts() {
	/* slice1 := make([]int, 5, 10)
	slice1[1] = 1
	slice1 = {1, 2, 3}
	fmt.Println("sli", slice1) */

	/* slice := []int{1, 2, 3}
	fmt.Println("Before append:", slice, "Capacity:", cap(slice))

	slice = append(slice, 4, 5, 6, 7)
	fmt.Println("After append:", slice, "Capacity:", cap(slice)) */

	/* slice := []int{1}
	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Println("Len:", len(slice), "Cap:", cap(slice))
	} */
	/* slice := []int{1}
	slice2 := slice
	slice2[0] = 11
	fmt.Println("", slice, slice2)
	*/
}
