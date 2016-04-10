apigen
==========
[![GoDoc](https://godoc.org/github.com/josephspurrier/apigen?status.svg)](https://godoc.org/github.com/josephspurrier/apigen)

Code Generator in Go

Package uses the ast standard library to generate Go code from templates. The package API may change.

The package can do the following things:
* Load a file
* Create a new file
* Copy a struct (with field comments and tags) to another package
* Change the package name
* Add imports
* Create a new imports section
* Print all comments in a package

Below is sample code to generate the code for a hello world application.

```go
var err error
var arrImports []string

arrImports = append(arrImports, "fmt")

// Create the package
gt := tree.New("main")
gt.AddImportSection(arrImports)
gt.AddHelloMainFunc()

// Write to a file
err = gt.WriteFile("output/hello/hello.go", true, 0644, 0644)
if err != nil {
	log.Println(err)
	return
}
```

Here is the output:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}
```

## Issues

When copying structs, the comments above structs do not copy. I believe this is related to the way ast was first written because it was a very early package. Links to the comment issues: [11880](https://github.com/golang/go/issues/11880) and [14629](https://github.com/golang/go/issues/14629)