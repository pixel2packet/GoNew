package main

import (
	"fmt"
)

func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// type stack[T any] struct { //LIFO (last in, First out)
// 	elements []T
// }

func main() {

	// myStack := stack[string]{
	// 	elements: []string{"go", "javascript"},
	// }
	// nums := []int{1, 2, 3}
	// names := []string{"go", "javascript"}
	vals := []bool{true, false, true}
	// printStringSlice(names)

	printSlice(vals, "john")
	// fmt.Println(myStack)

}
