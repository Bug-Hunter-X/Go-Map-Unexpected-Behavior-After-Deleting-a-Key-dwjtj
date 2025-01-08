```go
func main() {
    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2

    v, ok := m["c"]
    fmt.Println(v, ok) // Output: 0 false

    v, ok = m["a"]
    fmt.Println(v, ok) // Output: 1 true

    delete(m, "a")
    v, ok = m["a"]
    fmt.Println(v, ok) // Output: 0 false

    // the bug is here 
    v, ok = m["a"]
    fmt.Println(v, ok) // Output: 0 false
    m["a"] = 10 // this will overwrite previous value, but the key "a" is already deleted, 
                 // if this will be called after delete, it should return zero value with false
    fmt.Println(m["a"]) // Output: 10
}
```