package main

import (
	"fmt"
	"golang/module1"
)

func main() {
	fmt.Println(module1.IsValid("()"))    // true
	fmt.Println(module1.IsValid("()[]{")) // true
}
