package main

import (
	"log"

	"github.com/josephspurrier/apigen/tree"

	"modelgen"
)

// Source: http://goast.yuroyoro.net/
// Source: http://pastie.org/1574951#38,42
// Source: http://stackoverflow.com/questions/27976493/golang-static-identifier-resolution
// Source: https://golang.org/ref/spec

func main() {
	generateModel()
}

func generateModel() {
	var err error

	// Create the package
	gt := modelgen.New("model")

	// Imports
	i := gt.Import()
	i.AddImport("log")
	i.AddImport("fmt")
	gt.AddImport(i)

	// Main func
	f := gt.FuncDecl("main")
	f.AddComment("// Comment line 1")
	f.AddCallExpr("fmt.Print", "hello world")
	gt.AddFunc(f)

	// Test function
	g := gt.FuncDecl("testFunc").AddParam("monkey", "string").AddResult("", "error").AddReturnStmt(nil)
	gt.AddFunc(g)

	/*
		// Main func
		f := gtt.Func("main")
		f.AddComment("// Comment line 1")
		f.Comment("// Comment line 2")
		f.Param("varName", "string")
		f.Result("", "bool")
		f.Result("", "error")
		f.AddCallExpr("fmt.Print", 45)
		gtt.AddFunc(f)
	*/

	// Write to a file
	err = gt.WriteFile("output/model/model.go", true, 0644, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func generateHello() {
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
}
