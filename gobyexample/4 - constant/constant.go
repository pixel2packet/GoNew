// Constants
// Go supports _constants_ of character, string, boolean, and numeric values.

package main

import (
	"fmt"
	"math"
)

// // `const` declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s) // constant

	// // A `const` statement can appear anywhere a `var`
	// // statement can.
	const n = 50000000

	// // Constant expressions perform arithmetic with
	// // arbitrary precision.
	const d = 2e11 / n
	fmt.Println(d) // 4000
	fmt.Println(n) // 50000000

	// // A numeric constant has no type until it's given
	// // one, such as by an explicit conversion.
	fmt.Println(int64(d)) // 4000

	// // A number can be given a type by using it in a
	// // context that requires one, such as a variable
	// // assignment or function call. For example, here
	// // `math.Sin` expects a `float64`.
	fmt.Println(math.Sin(n)) // 0.8256467432733234

}
