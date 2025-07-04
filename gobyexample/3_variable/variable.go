package main

import "fmt"

func main() {

	// Variables
	// (variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.)

	// // `var` declares 1 or more variables.
	a := "initial"
	fmt.Println(a) // initial

	// // You can declare multiple variables at once.
	b, c := 2, 3
	fmt.Println(b, c) // 2 3

	// // Go will infer the type of initialized variables.
	d := true
	fmt.Println(d, !d) // true false

	// // Variables declared without a corresponding
	// // initialization are _zero-valued_. For example, the
	// // zero value for an `int` is `0`.
	var e int
	fmt.Println(e) // 0

	var f complex128
	fmt.Println(f) // (0+0i)

	var g bool
	fmt.Println(g) //false

}
