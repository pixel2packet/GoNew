# 🐹 Go Variadic Functions Practice

A Go program demonstrating **variadic functions**, which can accept a variable number of arguments — like `fmt.Println()`.

---

# 🚀 Getting Started

This example shows:
- How to define a function that takes any number of arguments
- How to loop over those arguments using `range`
- How to pass a slice into a variadic function using `...`

---

# 🔧 Topics Covered / Learnt

- Declaring a variadic function: `func sum(nums ...int)`
- Inside the function, `nums` is treated as a slice (`[]int`)
- Iterating over the slice with `for _, num := range nums`
- Calculating the sum of arguments
- Calling the function with:
  - Individual arguments: `sum(1, 2, 3)`
  - A slice using spread syntax: `sum(slice...)`
