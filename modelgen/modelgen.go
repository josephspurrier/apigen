package modelgen

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

// Var represents a variable so it's not confused with a string
type Var struct {
	Name string
}

// Variable returns a variable type
func Variable(v string) Var {
	return Var{
		Name: v,
	}
}

// Tree represents a fileset of code
type Tree struct {
	FileSet    *token.FileSet
	File       *ast.File
	CommentMap ast.CommentMap
}

// New creates a new Go package tree
func New(name string) *Tree {
	gt := &Tree{
		FileSet: token.NewFileSet(),
		File: &ast.File{
			Name: &ast.Ident{
				Name: name,
			},
			Doc: &ast.CommentGroup{
				List: []*ast.Comment{},
			},
		},
	}

	return gt
}

func ParseFile(code string) (*Tree, error) {
	gt := &Tree{}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		return gt, err
	}

	gt.FileSet = fset
	gt.File = f

	return gt, nil
}

// This adds a comment inbetween the package and name declaration so it doesn't work
// Couldn't figure it out
/*func (t *Tree) AddComment(comment string) {
	c := &ast.Comment{
		Text:  comment,
		Slash: 1,
	}
	t.File.Package = c.End() + 1
	t.File.Name.NamePos = c.End() + 7 + 1

	t.File.Doc.List = append(t.File.Doc.List, c)
}*/

// Bytes returns the code as a byte array
func (gt *Tree) Bytes(doFormat bool) ([]byte, error) {
	var output []byte
	var err error

	buffer := bytes.NewBuffer(output)
	if err = printer.Fprint(buffer, gt.FileSet, gt.File); err != nil {
		return nil, err
	}

	output = buffer.Bytes()

	if doFormat {
		// Format the source
		output, err = format.Source(output)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

// WriteFile writes the code to a file and create the folder structure
func (gt *Tree) WriteFile(filepath string, doFormat bool, dirPerm os.FileMode, filePerm os.FileMode) error {
	var err error

	// Create folders
	err = os.MkdirAll(path.Dir(filepath), dirPerm)
	if err != nil {
		return err
	}

	// Generate the code
	data, err := gt.Bytes(doFormat)
	if err != nil {
		return err
	}

	// Write file
	err = ioutil.WriteFile(filepath, data, filePerm)
	if err != nil {
		return err
	}

	return nil
}

// AddHelloMainFunc adds a main func the outputs: hello world
func (gt *Tree) AddMainFunc() {
	fd := &ast.FuncDecl{
		Name: &ast.Ident{
			Name: "main2",
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: []*ast.Ident{
							&ast.Ident{
								Name: "monkey",
							},
						},
						Type: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "string",
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Type: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "error",
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.Ident{
							Name: "fmt.Println",
						},
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: strconv.Quote("hello world"),
							},
						},
					},
				},
			},
		},
	}

	gt.File.Decls = append(gt.File.Decls, fd)
}
