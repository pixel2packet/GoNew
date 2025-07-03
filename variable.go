package main

import "fmt"

func main() {

	// Variables
	// (variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.)

	a := "initial"
	fmt.Println(a) // initial

	b, c := 2, 3
	fmt.Println(b, c) // 2 3

	d := true
	fmt.Println(d, !d) // true false

	var e int
	fmt.Println(e) // 0

	var f complex128
	fmt.Println(f) // (0+0i)

	var g bool
	fmt.Println(g) //false

}
