package main

import "fmt"

var foo int
var i, j int = 1, 2

func main() {
	foo = 55
	fmt.Println(foo)

	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)

	k := 3
	fmt.Println(i, j, k)
}
