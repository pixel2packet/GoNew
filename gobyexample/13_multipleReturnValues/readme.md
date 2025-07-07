# 🐹 Go Multiple Return Values Practice

A Go program demonstrating how a function can return **multiple values**. This is often used in Go for returning a result along with an error.

---

# 🚀 Getting Started

This example shows:
- How to define a function that returns more than one value
- How to assign multiple values from a function call
- How to ignore a return value using the blank identifier `_`

---

# 🔧 Topics Covered / Learnt

- Function returning multiple values: `func vals() (int, int)`
- Assigning returned values: `a, b := vals()`
- Ignoring one return value: `_, c := vals()`
- Useful in real-world situations like:  
  `result, err := doSomething()`
