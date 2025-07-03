package main

import "fmt"

func main() {
	// for loop
	// for is Go’s only looping construct. Here are some basic types of for loops.

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println(" ")

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println(" ")

	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println(" ")

	for {
		fmt.Println("loop")
		break
	}

	fmt.Println(" ")

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
