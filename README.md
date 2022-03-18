# genq

A Generic Queue package for Go.

```go
func main() {
    q := q.New[int]()
    fmt.Printf("messages: %v, want: %v\n", q.Len(), 0)
    q.Enqueue(3)
    q.Enqueue(2)
    q.Enqueue(90)
    fmt.Printf("messages: %v, want: %v\n", q.Len(), 3)
    for i := q.Len(); i > 0; i-- {
        m := q.Dequeue()
        fmt.Printf("messages: %v, want: %v\n", q.Len(), i)
        fmt.Printf("Value: %v\n", m.Value)
    }
}
```
