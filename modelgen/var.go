package modelgen

import (
	"go/ast"
	"go/token"
	"log"
)

// Var creates a var section
func (gt *Tree) Var() *GenDecl {
	gd := &ast.GenDecl{
		Lparen: 1,
		//Rparen: 1, // This can affect where the parens are and where the comments end up
		Tok:   token.VAR,
		Specs: []ast.Spec{},
	}

	return &GenDecl{gd}
}

// AddVar adds a var
func (gt *Tree) AddNewVar(varName, varType, varComment string) *Tree {
	for _, v := range gt.File.Decls {
		switch n := v.(type) {
		case *ast.GenDecl:
			if n.Tok == token.VAR {

				newSpecs := []ast.Spec{}

				// Loop through the old specs
				for _, s := range n.Specs {
					st := s.(*ast.ValueSpec)
					log.Println("Name:", st.Names[0].Name, st.Comment.List[0].Text)
					newSpecs = append(newSpecs, &ast.ValueSpec{
						Names: []*ast.Ident{
							&ast.Ident{
								Name: st.Names[0].Name,
							},
						},
						/*Type: &ast.Ident{
							Name: st.Type.(*ast.SelectorExpr).,
						},*/
						Comment: &ast.CommentGroup{
							List: []*ast.Comment{
								&ast.Comment{
									Text: st.Comment.List[0].Text,
								},
							},
						},
					})
				}

				// Add the new one
				newSpecs = append(newSpecs, &ast.ValueSpec{
					Names: []*ast.Ident{
						&ast.Ident{
							Name: varName,
						},
					},
					Type: &ast.Ident{
						Name: varType,
					},
					Comment: &ast.CommentGroup{
						List: []*ast.Comment{
							&ast.Comment{
								Text: varComment,
							},
						},
					},
				})

				n.Specs = newSpecs

				/*n.Specs = append(n.Specs, &ast.ValueSpec{
					Names: names,
					Type: &ast.Ident{
						Name: varType,
					},
					Doc: cg,
				})*/

				//gt.File.Comments = append(gt.File.Comments, cg)
			}
		default:
			//fmt.Printf("Missed Type: %#v\n", v)
		}
	}

	return gt
}

// AddVar adds a var
func (g *GenDecl) AddVar(varName, varType, varComment string) *GenDecl {
	g.GD.Specs = append(g.GD.Specs, &ast.ValueSpec{
		Names: []*ast.Ident{
			&ast.Ident{
				Name: varName,
			},
		},
		Type: &ast.Ident{
			Name: varType,
		},
		Comment: &ast.CommentGroup{
			List: []*ast.Comment{
				&ast.Comment{
					Text: varComment,
				},
			},
		},
	})

	return g
}

// AddVar adds a var section
func (gt *Tree) AddVar(gd *GenDecl) {
	gt.File.Decls = append(gt.File.Decls, gd.GD)
}
