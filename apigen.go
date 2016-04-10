package main

import (
	"go/parser"
	"log"

	"github.com/josephspurrier/apigen/tree"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)
}

func main() {
	generateHello()
	generateModel()
}

// Generate a package that outputs: hello world
func generateHello() {
	var err error
	var arrImports []string

	packageName := "main"
	arrImports = append(arrImports, "fmt")

	// Create the package
	gt := tree.New(packageName)
	gt.AddImportSection(arrImports)
	gt.AddHelloMainFunc()

	// Write to a file
	err = gt.WriteFile("output/hello/hello.go", true, 0644, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

// Generate a model from a template and a spec
func generateModel() {

	var err error
	var arrImports []string
	gt := &tree.Tree{}
	spec := &tree.Tree{}

	entityName := "user"

	// Read the file
	gt, err = tree.Load("template/model.go", 0)
	if err != nil {
		log.Println(err)
		return
	}

	// Set the package name
	gt.SetPackageName(entityName)

	// Update the variable value
	err = gt.ChangeConstString("bucketName", entityName)
	if err != nil {
		log.Println(err)
		return
	}

	// Read the spec
	spec, err = tree.Load("spec/user/user.go", parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}

	// Copy the struct fields
	structFields, structComments, err := spec.StructFields("Entity")
	if err != nil {
		log.Println(err)
		return
	}

	// Replace the struct
	err = gt.ChangeStruct("Entity", structFields, structComments)
	if err != nil {
		log.Println(err)
		return
	}

	// Add the imports
	for i := 0; i < len(arrImports); i++ {
		gt.AddImport(arrImports[i])
	}

	// Write to a file
	err = gt.WriteFile("output/user/user.go", true, 0644, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}
