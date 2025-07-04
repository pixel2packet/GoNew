// _Slices_ are an important data type in Go, giving
// a more powerful interface to sequences than arrays.

package main

import (
	"fmt"
	"slices"
)

func main() {

	// Unlike arrays, slices are typed only by the elements they
	// contain (not the number of elements).
	// An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("unintit:", s, s == nil, len(s) == 0) // unintit: [] true true

	// To create a slice with non-zero length, use the builtin `make`.
	// Here we make a slice of `string`s of length `3` (initially zero-valued).
	// By default a new slice's capacity is equal to its length; if we know the slice
	// is going to grow ahead of time, it's possible to pass a capacity explicitly
	// as an additional parameter to `make`.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // emp: [  ] len: 3 cap: 3

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)    // [a b c]
	fmt.Println("get:", s[2]) // c

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s)) // 3

	// Extra operations, 1st is the builtin append,
	// which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) // [a b c d e f]

	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) // [a b c d e f]

	// Slices support a “slice” operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("cpy:", l) // [c d e]

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2:", l) // [a b c d e]

	// And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3:", l) // [c d e f]

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // [g h i]

	// The slices package contains a number of useful utility functions for slices.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2") // t == t2
	}

	// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	// twoD := make([][]int, 3)
	// for i := range 3 {
	// 	innerLen := i + 1
	// 	twoD[i] = make([]int, innerLen)
	// 	for j := range innerLen {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2D:", twoD) // 2D: [[0] [1 2] [2 3 4]]

	twoD := make([][]int, 3)
	for i := range 3 {
		innerlane := i + 1
		twoD[i] = make([]int, innerlane)
		for j := range innerlane {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)
}
