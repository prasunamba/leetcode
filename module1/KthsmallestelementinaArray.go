package module1

import (
	"fmt"
	"sort"
)

func KthsmallestelementinaArray() {
	nums := []int{2, 0, 11, 15}
	sort.Ints(nums)
	k := 3
	fmt.Println("nums", nums)
	fmt.Println("", nums[k-1])
}
