package main

import (
	"fmt"
	"math"
)

// // Constants
// // (Go supports constants of character, string, boolean, and numeric values.)
const s string = "constant"

func main() {
	fmt.Println(s) // constant

	const n = 50000000

	const d = 2e11 / n
	fmt.Println(d) // 4000
	fmt.Println(n) // 50000000

	fmt.Println(int64(d)) // 4000

	fmt.Println(math.Sin(n)) // 0.8256467432733234

}
