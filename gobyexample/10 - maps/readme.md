# 🐹 Go Maps Practice

A Go program demonstrating how to declare, use, compare, and manipulate maps (key-value data structures).

---

# 🚀 Getting Started

This example helps understand:
- Creating and updating maps
- Accessing and checking keys
- Deleting and clearing map entries
- Using `maps.Equal()` to compare maps
- Storing values from maps into arrays

---

# 🔧 Topics Covered / Learnt

- Creating an empty map with `make(map[key]value)`
- Adding key-value pairs: `m["key"] = value`
- Accessing values using `m["key"]`
- Accessing a missing key returns the zero value (e.g. `0`)
- Checking if a key exists using:  
  `_, exists := m["key"]`
- Deleting a key using `delete(m, "key")`
- Clearing all entries with `clear(m)`
- Comparing two maps with `maps.Equal(map1, map2)`
- Creating a map with literal syntax:  
  `map[string]int{"foo": 1, "bar": 2}`
- Inserting map values into an array:  
  `j := [...]int{v1, v3, v2}`