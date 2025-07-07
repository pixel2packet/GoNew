// _Slices_ are an important data type in Go, giving
// a more powerful interface to sequences than arrays.

package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("unintit:", s, s == nil, len(s) == 0) // unintit: [] true true

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // emp: [  ] len: 3 cap: 3

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)    // [a b c]
	fmt.Println("get:", s[2]) // c

	fmt.Println("len:", len(s)) // 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) // [a b c d e f]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) // [a b c d e f]

	l := s[2:5]
	fmt.Println("cpy:", l) // [c d e]

	l = s[:5]
	fmt.Println("sl2:", l) // [a b c d e]

	l = s[2:]
	fmt.Println("sl3:", l) // [c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // [g h i]

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2") // t == t2
	}

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
