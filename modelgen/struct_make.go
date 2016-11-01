package modelgen

import (
	"errors"
	"go/ast"
)

var (
	ErrVarMustString = errors.New("variable type must be a string")
	ErrVarNotFound   = errors.New("variable not found")
)

type StructChanger struct {
	Name     string
	Fields   []*ast.Field
	Comments *ast.CommentGroup
	Ok       bool
}

// ChangeStruct changes the struct field list
func (gt *Tree) ChangeStruct(varName string, varValue []*ast.Field, varComments *ast.CommentGroup) error {
	sc := &StructChanger{
		Name:     varName,
		Fields:   varValue,
		Comments: varComments,
	}

	ast.Walk(sc, gt.File)

	if !sc.Ok {
		return ErrVarNotFound
	}

	return nil
}

// Visit walks the tree for StructChanger
func (v *StructChanger) Visit(n ast.Node) (w ast.Visitor) {
	switch spec := n.(type) {
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
	return v
}
