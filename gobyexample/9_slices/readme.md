# 🐹 Go Slices Practice

A Go program demonstrating the use of **slices**, which provide a more flexible and powerful way to handle sequences compared to arrays.

---

# 🚀 Getting Started

This example covers:
- Declaring and initializing slices
- Appending and copying elements
- Slicing with ranges
- Comparing slices
- Creating 2D slices dynamically

---

# 🔧 Topics Covered / Learnt

- Declaring a nil slice (`var s []string`)
- Initializing with `make([]T, length)`
- Setting and getting values: `s[i] = value`, `s[i]`
- Checking length with `len()` and capacity with `cap()`
- Appending elements using `append(...)`
- Copying slices with `copy(dst, src)`
- Slicing with ranges: `s[low:high]`, `s[:n]`, `s[m:]`
- Declaring and comparing slices: `[]T{...}` + `slices.Equal(a, b)`
- Creating and filling 2D slices dynamically (`[][]int`)