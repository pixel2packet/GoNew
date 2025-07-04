// for loop
// for is Go’s only looping construct. Here are some basic types of for loops.

package main

import "fmt"

func main() {

	// // The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println(" ")

	// // A classic initial/condition/after `for` loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println(" ")

	// // Another way of accomplishing the basic "do this
	// // N times" iteration is `range` over an integer.
	for i := range 4 {
		fmt.Println("range", i)
	}

	fmt.Println(" ")

	// // `for` without a condition will loop repeatedly
	// // until you `break` out of the loop or `return` from
	// // the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println(" ")

	// // You can also `continue` to the next iteration of
	// the loop.
	for n := range 8 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// Answers
// 1
// 2
// 3

// 0
// 1
// 2

// range 0
// range 1
// range 2
// range 3

// loop

// 1
// 3
// 5
// 7
