// In Go, an _array_ is a numbered sequence of elements of a specific length.
// In typical Go code, [slices](slices) are much more common;
// arrays are useful in some special scenarios.

package main

import "fmt"

func main() {

	// Here we create an array `a` that will hold exactly 5 `int`s.
	// The type of elements and length are both part of the array's type.
	// By default an array is zero-valued, which for `int`s means `0`s.
	var a [5]int
	fmt.Println("emp:", a) //emp = empty //[0 0 0 0 0]

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.
	a[3] = 100                      //array[index] = value
	fmt.Println("set:", a)          //Updated array with a[4] = 100
	fmt.Println("get:", a[3], a[4]) //array[index]

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array
	// in one line.
	b := [5]int{2, 4, 6, 8, 10}
	fmt.Println("dec:", b)

	// You can also have the compiler count the number of
	// elements for you with `...`
	b = [...]int{2, 4, 6, 8, 10}
	fmt.Println("dec:", b)

	// If you specify the index with `:`, the elements in
	// between will be zeroed.
	b = [...]int{100, 3: 300, 400}
	fmt.Println("idx:", b) //idx = index

	// Array types are one-dimensional, but you can compose
	// types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)

	// You can create and initialize multi-dimensional
	// arrays at once too.
	twoD = [2][3]int{
		{10, 20, 30},
		{20, 30, 40},
	}
	fmt.Println("2D: ", twoD)

}
