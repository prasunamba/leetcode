/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/

package main

import "fmt"

type mystructt struct {
	name string
	age  int
}

func Swap() {
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println("before ", a, b, c)
	a, b, c = b, c, a
	fmt.Println("after ", a, b, c)
}
