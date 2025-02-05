/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/

package main

import "fmt"

type shape interface {
	area() int
}
type rectangle struct {
	lenth, breadth int
}

func (r rectangle) area() int {
	return r.lenth * r.breadth
}
func calculate(s shape) {
	fmt.Println(s.area())
}
func inter() {
	fmt.Println("Hello World")
	r1 := rectangle{
		lenth:   10,
		breadth: 5,
	}
	calculate(r1)
}
