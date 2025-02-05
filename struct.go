/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/

package main

import "fmt"

type mystruct struct {
	name string
	age  int
}

func structexample() {
	fmt.Println("Hello World")
	p1 := mystruct{
		name: "gopi",
		age:  32,
	}
	fmt.Println("p1", p1)
	fmt.Println("name ", p1.name)
	fmt.Println("age  ", p1.age)
	p1.age = 12
	fmt.Println("p1", p1)
}
