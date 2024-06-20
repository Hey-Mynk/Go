# Notes on the `fmt` Package in Go

The `fmt` package in Go provides functionality for formatted I/O operations. It is commonly used for printing output to the console or formatting strings.

## Printing to the Console

The `fmt` package provides the `Print`, `Printf`, and `Println` functions for printing output to the console. These functions accept a variety of arguments and format them according to the provided format string.

Example usage:

```go
package main

import "fmt"

func main() {
    fmt.Print("Hello, ")
    fmt.Println("world!")
    fmt.Printf("The value of pi is approximately %.2f\n", 3.14159)
}
```

Output:

```
Hello, world!
The value of pi is approximately 3.14
```

## Formatting Strings

The `fmt` package also provides functions for formatting strings, such as `Sprintf` and `Fprintf`. These functions allow you to format strings without printing them to the console.

Example usage:

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 30
    formattedString := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
    fmt.Println(formattedString)
}
```

Output:

```
My name is Alice and I am 30 years old.
```

## Formatting Options

The `fmt` package supports a variety of formatting options, such as specifying the width and precision of numeric values, padding strings, and aligning text. These options can be specified using verbs and flags in the format string.

Example usage:

```go
package main

import "fmt"

func main() {
    width := 8
    precision := 2
    value := 123.456789
    fmt.Printf("Value: %*.*f\n", width, precision, value)
}
```

Output:

```
Value:   123.46
```


### func FileServer(root FileSystem) Handler


FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.

As a special case, the returned file server redirects any request ending in "/index.html" to the same path, without the final "index.html".

To use the operating system's file system implementation, use http.Dir: