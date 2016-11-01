package modelgen

import (
	"fmt"
	"go/ast"
)

type VarWrite struct {
	Comments *ast.CommentGroup
	Expr     *ast.Expr
	Names    []*ast.Ident

	Ok bool
}

// ChangeStruct changes the struct field list
func (gt *Tree) WriteVar(sc VarRead) error {
	sr := &VarWrite{
		Names:    sc.Names,
		Expr:     sc.Expr,
		Comments: sc.Comments,
	}

	ast.Walk(sr, gt.File)

	if !sc.Ok {
		return ErrVarNotFound
	}

	return nil
}

// Visit walks the tree
func (v *VarWrite) Visit(n ast.Node) (w ast.Visitor) {
	n = DiscoverNode(n)
	return v

	/*switch t := n.(type) {
	case *ast.GenDecl:
		if t.Tok == token.VAR {

			t.Specs = []ast.Spec{}

			t.Specs = append(t.Specs, &ast.ValueSpec{
				Names:   v.Names,
				Comment: v.Comments,
				Type:    *v.Expr,
			})

			for _, spec := range t.Specs {
				if vs, ok := spec.(*ast.ValueSpec); ok {

					vs.Names = v.Names
					vs.Comment = v.Comments
					vs.Type = *v.Expr
					return nil
				}
			}
		}
		case *ast.TypeSpec:
		if spec.Name.String() == v.Name {
			switch structType := spec.Type.(type) {
			case *ast.StructType:
				structType.Fields.List = DiscoverList(v.Fields)
				spec.Comment = v.Comments
				v.Ok = true
				return nil
			}
		}
	}
	return v*/
}

// DiscoverType builds clean fields for ast
func DiscoverNode(node ast.Node) ast.Node {
	switch t := node.(type) {
	/*case *ast.File:
		return &ast.File{
			Name: t.Name,
			Doc: &ast.CommentGroup{
				List: t.Doc.List,
			},
		}
	case *ast.Comment:
		return &ast.Comment{
			Text: t.Text,
		}*/
	/*case *ast.Ident:
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
		}*/
	default:
		fmt.Printf("Missed Type: %#v\n", t)
	}
	return nil
}
