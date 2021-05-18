# fmt

A drop-in replacement package for the Go stdlib `fmt` package that provides
three additional functions that behave the same as their counterpart `printf`
functions, but append a new line on the end.

```go
func Printfln(format string, a ...interface{}) (n int, err error)
func Fprintfln(w io.Writer, format string, a ...interface{})
func Sprintfln(format string, a ...interface{})
```

This package is an implementation of a proposal for the Go stdlib `fmt` package:
https://github.com/golang/go/issues/46190

## Usage

```go
import "4d63.com/fmt"
```

```go
fmt.Printfln("Hello %s!", "world")
fmt.Printfln("G'day %s!", "world")
// Output:
// Hello world!
// G'day world!
```
