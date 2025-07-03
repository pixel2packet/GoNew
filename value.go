package main

import "fmt"

func main() {

	// Values
	// (Go has various value types including strings, integers, floats, booleans, etc.)

	fmt.Println("go" + "lang") // golang

	fmt.Println("1+1 =", 1+1)         //1+1 = 2
	fmt.Println("7.0/3.0 =", 7.0/3.0) //7.0/3.0 = 2.3333333333333335

	fmt.Println("T && T:", true && true)   // T && T: true
	fmt.Println("T && F:", true && false)  // T && F: false
	fmt.Println("F && T:", false && true)  // F && T: false
	fmt.Println("F && F:", false && false) // F && F: false

	fmt.Println("T || T:", true || true)   // T || T: true
	fmt.Println("T || F:", true || false)  // T || F: true
	fmt.Println("F || T:", false || true)  // F || T: true
	fmt.Println("F || F:", false || false) // F || F: false

	fmt.Println(!true)            // false
	fmt.Println(!false && !false) // true
	fmt.Println(!false || !false) // true
	fmt.Println(!true && !true)   // false
	fmt.Println(!true || !true)   // false
	fmt.Println(!!true)           // true
}
