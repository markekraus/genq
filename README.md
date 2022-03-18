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
}
```
