package modelgen

import (
	"go/ast"
	"go/token"
	"strconv"
)

type GenDecl struct {
	GD *ast.GenDecl
}

// Import creates an import section
func (gt *Tree) Import() *GenDecl {
	gd := &ast.GenDecl{
		Lparen: 1,
		Rparen: 1,
		Tok:    token.IMPORT,
		Specs:  []ast.Spec{},
	}

	return &GenDecl{gd}
}

// AddNewImport adds an import
func (gt *Tree) AddNewImport(imp string) *Tree {
	for _, v := range gt.File.Decls {
		switch n := v.(type) {
		case *ast.GenDecl:
			if n.Tok == token.IMPORT {
				n.Specs = append(n.Specs, &ast.ImportSpec{
					Path: &ast.BasicLit{
						Value: strconv.Quote(imp),
					},
				})
			}
		default:
			//fmt.Printf("Missed Type: %#v\n", v)
		}
	}

	return gt
}

// AddImport adds an import
func (g *GenDecl) AddImport(imp string) *GenDecl {
	g.GD.Specs = append(g.GD.Specs, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Value: strconv.Quote(imp),
		},
	})

	return g
}

// AddImport adds an import section with import paths
func (gt *Tree) AddImport(gd *GenDecl) {
	gt.File.Decls = append(gt.File.Decls, gd.GD)
}

// AddImport adds an import section with import paths
func (gt *Tree) AddImportSection(imports []string) {
	var specs []ast.Spec

	// Add all the imports
	for i := 0; i < len(imports); i++ {
		specs = append(specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Value: strconv.Quote(imports[i]),
			},
		})
	}

	// Generate the import declaration
	decl := &ast.GenDecl{
		Lparen: 1,
		Rparen: 1,
		Tok:    token.IMPORT,
		Specs:  specs,
	}
	gt.File.Decls = append(gt.File.Decls, decl)
}
