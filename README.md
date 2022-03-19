# genq

A Generic Queue (FIFO/LILO) package for Go.

```bash
go get github.com/markekraus/genq
```

```go
package main

import (
    "fmt"
    genq "github.com/markekraus/genq/pkg"
)

type mytype struct {
    a, b int
}

func main() {
    
}

func main() {
    q := genq.New[int]()
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

    q2 := genq.New[*mytype]()
    m1 := &mytype{1, 2}
    q2.Enqueue(&mytype{1, 2})
    m2 := q.Enqueue(m1).Value
    fmt.Printf("%v\n", m2 == m1)
    fmt.Printf("%v\n", m2.a == m1.a)
    fmt.Printf("%v\n", m2.b == m1.b)
}
```
