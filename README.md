infr
---

`Into` / `From` interface similar to `Into` / `From` trait implementation in Rust.

```go
import (
    "fmt"
    "github.com/fudanchii/infr"
)

// let's convert int to string with `From` / `Into` trait pattern
type Str string

// Implement `From` method, and we can use `Into` for free.
func (s Str) From(n int) Str {
    return Str(fmt.Sprintf("%d", n))
}

func main() {
    num := 42
    strnum := infr.Into[Str](num)
    fmt.Println("the answer to life, universe, and everything: %s", strnum)
}
```

licensed under:

2-clause BSD
