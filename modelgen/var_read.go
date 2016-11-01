package modelgen

import (
	"fmt"
	"go/ast"
	"go/token"
)

type VarRead struct {
	Comments *ast.CommentGroup
	Expr     *ast.Expr
	Names    []*ast.Ident

	Ok bool
}

// StructFields returns the struct fields
func (gt *Tree) ReadVar() (VarRead, error) {
	sc := &VarRead{}

	ast.Walk(sc, gt.File)

	if !sc.Ok {
		return *sc, ErrVarNotFound
	}

	return *sc, nil
}

// Visit walks the tree
func (v *VarRead) Visit(n ast.Node) (w ast.Visitor) {
	switch t := n.(type) {
	case *ast.GenDecl:
		if t.Tok == token.VAR {
			for _, spec := range t.Specs {
				if vs, ok := spec.(*ast.ValueSpec); ok {
					namesNew := []*ast.Ident{}

					for _, name := range vs.Names {
						namesNew = append(namesNew, &ast.Ident{
							Name: name.Name,
						})
					}

					tt := DiscoverType(vs.Type)

					commentsNew := []*ast.Comment{}

					for _, comment := range vs.Comment.List {
						commentsNew = append(commentsNew, &ast.Comment{
							Text: comment.Text,
						})
					}

					v.Expr = &tt
					v.Comments = &ast.CommentGroup{
						List: commentsNew,
					}
					v.Names = namesNew

					v.Ok = true
				}
			}
		}
	case *ast.TypeSpec:
		/*log.Println("Name:", spec.Name)
		if spec.Name.String() == v.Name {
			switch structType := spec.Type.(type) {
			case *ast.StructType:
				v.Fields = structType.Fields.List
				v.Comments = spec.Comment
				v.Ok = true
				return nil
			}
		}*/
	default:
		//fmt.Printf("Missed Type: %#v\n", n)
	}

	return v
}

// DiscoverType builds clean fields for ast
func DiscoverType(express ast.Expr) ast.Expr {
	switch t := express.(type) {
	case *ast.Ident:
		return &ast.Ident{
			Name: t.Name,
		}
	case *ast.StarExpr:
		return &ast.StarExpr{
			X: DiscoverType(t.X),
		}
	case *ast.ArrayType:
		return &ast.ArrayType{
			Elt: DiscoverType(t.Elt),
		}
	case *ast.MapType:
		return &ast.MapType{
			Key:   t.Key,
			Value: t.Value,
		}
	case *ast.InterfaceType:
		return &ast.InterfaceType{
			Methods: t.Methods,
		}
	case *ast.SelectorExpr:
		return &ast.SelectorExpr{
			X: DiscoverType(t.X),
			Sel: &ast.Ident{
				Name: t.Sel.Name,
			},
		}
	case *ast.StructType:
		return &ast.StructType{
			Fields: &ast.FieldList{
				List: DiscoverList(t.Fields.List),
			},
		}
	case *ast.ChanType:
		return &ast.ChanType{
			Arrow: t.Arrow,
			Dir:   t.Dir,
			Value: DiscoverType(t.Value),
		}
	case *ast.FuncType:
		return &ast.FuncType{
			Params: &ast.FieldList{
				List: DiscoverList(t.Params.List),
			},
			Results: &ast.FieldList{
				List: DiscoverList(t.Results.List),
			},
		}
	default:
		fmt.Printf("Missed Type: %#v\n", express)
	}
	return nil
}

// DiscoverList builds clean lists for ast
func DiscoverList(list []*ast.Field) []*ast.Field {
	var l []*ast.Field

	for i := 0; i < len(list); i++ {
		field := &ast.Field{}

		// Comments
		if list[i].Comment != nil {

			// Create a comment group
			field.Comment = &ast.CommentGroup{}

			// Loop through the comments
			for c := 0; c < len(list[i].Comment.List); c++ {
				field.Comment.List = append(field.Comment.List, &ast.Comment{
					Text: list[i].Comment.List[c].Text,
				})
			}

		}

		// Loop through the names
		for n := 0; n < len(list[i].Names); n++ {
			field.Names = append(field.Names, &ast.Ident{
				Name: list[i].Names[n].Name,
			})
		}

		// Tags
		if list[i].Tag != nil {
			field.Tag = &ast.BasicLit{
				Kind:  list[i].Tag.Kind,
				Value: list[i].Tag.Value,
			}
		}

		// Types
		field.Type = DiscoverType(list[i].Type)

		l = append(l, field)
	}

	return l
}
