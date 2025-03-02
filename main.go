package main

import (
	"fmt"
	"golang/module1"
)

func main() {
	defer fmt.Println("main exit")
	module1.CharFrequencyOptimized()
}
