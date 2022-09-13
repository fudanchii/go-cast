infr
---

`Into` / `From` interface similar to [`Into`](https://doc.rust-lang.org/std/convert/trait.Into.html) / [`From`](https://doc.rust-lang.org/std/convert/trait.From.html) trait implementation in Rust.

To quote the doc from the Rust version:

> One should always prefer implementing `From` over `Into` because implementing `From` automatically provides one with an implementation of `Into` thanks to the blanket implementation in the standard library.

That is, for the implementation, one should always prefer to implement `From`, and not `Into`.

> Prefer using `Into` over using `From` when specifying trait bounds on a generic function. This way, types that directly implement `Into` can be used as arguments as well.

As for specifying interface bound, e.g. parameter type, one should always use `Into`, and not `From`.
This is more cumbersome in Go, as Go doesn't have blanket implementation for interfaces, it's only possible by wrapping the type into another type (e.g. a struct), then write the `interface` implementation for this struct, acting as the blanket implementation for the other wrapped type.  

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

func StrPrinter(s infr.IntoType[Str]) string {
    return string(s.Into())
}

func main() {
    num := 42
    strnum := infr.Into[Str](num)
    fmt.Println("the answer to life, universe, and everything: %s", strnum)
    
    fmt.Println("alternatively")
    
    strnum1 := StrPrinter(infr.FI[Str]{num})
    fmt.Println("the answer to life, universe, and everything: %s", strnum1)
}
```

licensed under: 2-clause BSD
