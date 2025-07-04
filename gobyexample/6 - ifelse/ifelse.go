// If/else
// // Branching with if and else in Go is straight-forward.

package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// // You can have an `if` statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// // Logical operators like `&&` and `||` are often
	// // useful in conditions.
	if 6%3 == 0 && 6%1 == 0 {
		fmt.Println("8 is divisible by 3 and 1")
	}

	// // A statement can precede conditionals; any variables
	// // declared in this statement are available in the current
	// // and all subsequent branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Answers
// 7 is odd
// 8 is divisible by 4
// 8 is divisible by 3 and 1
// 9 has 1 digit
