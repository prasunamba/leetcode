/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/

package main

import (
	"fmt"
)

func fact(i int) int {
	if i < 2 {
		return 1
	}
	return fact(i-1) * i
}
func factorial() {

	result := fact(6)
	fmt.Println("result", result)
}
