package modelgen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
	"strconv"
)

func (gt *Tree) AddFunc(fd *FuncDecl) {
	gt.File.Decls = append(gt.File.Decls, fd.FD)
}

type FuncDecl struct {
	FD *ast.FuncDecl
}

// FuncDecl returns a function declaration
func (gt *Tree) FuncDecl(name string) *FuncDecl {
	fd := &ast.FuncDecl{
		Name: &ast.Ident{
			Name: name,
		},
		// Exclude to prevent empty parenthesis from showing up if not used
		/*Recv: &ast.FieldList{
			List: []*ast.Field{},
		},*/
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{},
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{},
		},
	}

	return &FuncDecl{fd}
}

func (f *FuncDecl) AddComment(comment string) *FuncDecl {
	f.FD.Doc.List = append(f.FD.Doc.List, &ast.Comment{
		Text: comment,
	})

	return f
}

func (f *FuncDecl) AddParam(varName, varType string) *FuncDecl {
	f.FD.Type.Params.List = append(f.FD.Type.Params.List, &ast.Field{
		Names: []*ast.Ident{
			&ast.Ident{
				Name: varName,
			},
		},
		Type: &ast.BasicLit{
			Kind:  token.STRING,
			Value: varType,
		},
	})

	return f
}

func (f *FuncDecl) AddResult(varName, varType string) *FuncDecl {
	f.FD.Type.Results.List = append(f.FD.Type.Results.List, &ast.Field{
		Names: []*ast.Ident{
			&ast.Ident{
				Name: varName,
			},
		},
		Type: &ast.BasicLit{
			Kind:  token.STRING,
			Value: varType,
		},
	})

	return f
}

func (f *FuncDecl) AddRecv(varName, varType string) *FuncDecl {
	f.FD.Recv = &ast.FieldList{
		List: []*ast.Field{
			&ast.Field{
				Names: []*ast.Ident{
					&ast.Ident{
						Name: varName,
					},
				},
				Type: &ast.BasicLit{
					Kind:  token.STRING,
					Value: varType,
				},
			},
		},
	}

	return f
}

func (f *FuncDecl) AddCallExpr(funcName string, funcArg ...interface{}) *FuncDecl {
	f.FD.Body.List = append(f.FD.Body.List, &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.Ident{
				Name: funcName,
			},
			Args: buildExprList(funcArg...),
		},
	})

	return f
}

func (f *FuncDecl) AddReturnStmt(funcArg ...interface{}) *FuncDecl {
	f.FD.Body.List = append(f.FD.Body.List, &ast.ReturnStmt{
		Results: buildExprList(funcArg...),
	})

	return f
}

func (f *FuncDecl) AddDeclStmt(varName, varType string) *FuncDecl {
	f.FD.Body.List = append(f.FD.Body.List, &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						&ast.Ident{
							Name: varName,
						},
					},
					Type: &ast.Ident{
						Name: varType,
					},
				},
			},
		},
	})

	return f
}

func Ident(names ...string) []ast.Expr {
	expList := []ast.Expr{}

	for _, arg := range names {
		expList = append(expList, &ast.Ident{
			Name: fmt.Sprint(arg),
		})
	}

	return expList
}

func (f *FuncDecl) AddAssignStmt(varName []ast.Expr, varToken token.Token, varValue interface{}) *FuncDecl {
	f.FD.Body.List = append(f.FD.Body.List, &ast.AssignStmt{
		Lhs: varName,
		Tok: varToken,
		Rhs: buildExprListIdentRhs(varValue),
	})

	return f
}

// AddNewImport adds an import
func (gt *Tree) AddNewAssignStmt(funcName string, varName []ast.Expr, varToken token.Token, varValue interface{}) *Tree {
	for _, v := range gt.File.Decls {
		switch n := v.(type) {
		case *ast.FuncDecl:
			if n.Name.Name == funcName {
				n.Body.List = append(n.Body.List, &ast.AssignStmt{
					Lhs: varName,
					Tok: varToken,
					Rhs: buildExprListIdentRhs(varValue),
				})
			}
		default:
			//fmt.Printf("Missed Type: %#v\n", v)
		}
	}

	return gt
}

func buildExprListIdent(funcArg []ast.Expr) []ast.Expr {
	expList := []ast.Expr{}

	for _, arg := range funcArg {
		expList = append(expList, &ast.Ident{
			Name: fmt.Sprint(arg),
		})

		/*if fmt.Sprint(arg) == "var" {
			expList = append(expList, &ast.Ident{
				Name: "joe",
			})
		}

		/*if arg == nil {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: "nil",
			})
		} else if reflect.TypeOf(arg) == reflect.TypeOf(Var{}) {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: arg.(Var).Name,
			})
		} else if reflect.TypeOf(arg).Kind() == reflect.String {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(fmt.Sprint(arg)),
			})
		} else {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprint(arg),
			})
		}*/
	}

	return expList
}

func buildExprListIdentRhs(funcArg ...interface{}) []ast.Expr {
	expList := []ast.Expr{}

	for _, arg := range funcArg {
		expList = append(expList, &ast.Ident{
			Name: fmt.Sprint(arg),
		})

		/*expList = append(expList, &ast.Ident{
			Name: "joe",
		})*/

		/*if arg == nil {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: "nil",
			})
		} else if reflect.TypeOf(arg) == reflect.TypeOf(Var{}) {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: arg.(Var).Name,
			})
		} else if reflect.TypeOf(arg).Kind() == reflect.String {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(fmt.Sprint(arg)),
			})
		} else {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprint(arg),
			})
		}*/
	}

	return expList
}

// *****************************************************************************
// Helpers
// *****************************************************************************

func buildExprList(funcArg ...interface{}) []ast.Expr {
	expList := []ast.Expr{}

	for _, arg := range funcArg {
		if arg == nil {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: "nil",
			})
		} else if reflect.TypeOf(arg) == reflect.TypeOf(Var{}) {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: arg.(Var).Name,
			})
		} else if reflect.TypeOf(arg).Kind() == reflect.String {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(fmt.Sprint(arg)),
			})
		} else {
			expList = append(expList, &ast.BasicLit{
				Kind:  token.STRING,
				Value: fmt.Sprint(arg),
			})
		}
	}

	return expList
}

// *****************************************************************************
// Only for testing
// *****************************************************************************

func Output2(code string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", `// Package comment
package main`, parser.ParseComments)
	if err != nil {
		log.Println(err)
	}
	//log.Printf("%#v", f.Decls[0])
	log.Println(f.Doc)
	log.Println(f)

}

func Output(code string) {
	//fset := token.NewFileSet()
	f, err := parser.ParseExpr(`test = string`)

	if err != nil {
		log.Println(err)
	}
	log.Println(f)

}

/*func CallExpr(name string, args ...interface{}) *ast.CallExpr {
	return &ast.CallExpr{
		Fun: &ast.Ident{
			Name: funcName,
		},
		Args: buildExprList(args...),
	}
}*/
