```go
func main() {
    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2

    v, ok := m["c"]
    fmt.Println(v, ok) // Output: 0 false

    v, ok = m["a"]
    fmt.Println(v, ok) // Output: 1 true

    deleted := delete(m, "a")
    fmt.Println(deleted) //Output: true
    v, ok = m["a"]
    fmt.Println(v, ok) // Output: 0 false

    // Check if the key exists before overwriting
    if _, ok := m["a"]; !ok {
        m["a"] = 10
    }
    fmt.Println(m["a"]) // Output: 10
}
```