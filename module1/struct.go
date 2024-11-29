package module1

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
