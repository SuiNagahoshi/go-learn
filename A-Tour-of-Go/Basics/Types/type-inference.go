package main

import "fmt"

func Inference() {
	var i = 55
	var f = 66.7
	var c = -5 + 12i
	var s = "Hello"

	fmt.Printf("%T %T %T %T\n", i, f, c, s)
}
