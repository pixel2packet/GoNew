package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	v2 := m["k2"]
	fmt.Println("v2:", v2)

	j := [...]int{v1, v3, v2}
	fmt.Println("mar:", j) // mar = mapArray, mar: [7 0 13]

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map-del:", m)

	clear(m)
	fmt.Println("map-clr:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	m["k1"] = 5

	_, exists := m["k1"]
	fmt.Println(exists)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
