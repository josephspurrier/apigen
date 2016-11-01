package modelgen

import "go/ast"

type StructCopier struct {
	Name     string
	Fields   []*ast.Field
	Comments *ast.CommentGroup
	Ok       bool
}

// StructFields returns the struct fields
func (gt *Tree) StructFields(varName string) ([]*ast.Field, *ast.CommentGroup, error) {
	sc := &StructCopier{
		Name: varName,
	}

	ast.Walk(sc, gt.File)

	if !sc.Ok {
		return nil, nil, ErrVarNotFound
	}

	return sc.Fields, sc.Comments, nil
}

// Visit walks the tree for StructCopier
func (v *StructCopier) Visit(n ast.Node) (w ast.Visitor) {
	switch spec := n.(type) {
	case *ast.TypeSpec:
		if spec.Name.String() == v.Name {
			switch structType := spec.Type.(type) {
			case *ast.StructType:
				v.Fields = structType.Fields.List
				v.Comments = spec.Comment
				v.Ok = true
				return nil
			}
		}
	}

	return v
}
