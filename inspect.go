package brenda

import (
	"fmt"
	"go/ast"
)

func _(node ast.Node) bool {
	if node == nil {
		return true
	}
	switch n := node.(type) {
	// Comments and fields
	case *ast.Comment:
	// nothing to do

	case *ast.CommentGroup:
	/*for _, c := range n.List {
		Walk(v, c)
	}*/

	case *ast.Field:
	/*if n.Doc != nil {
		Walk(v, n.Doc)
	}
	walkIdentList(v, n.Names)
	Walk(v, n.Type)
	if n.Tag != nil {
		Walk(v, n.Tag)
	}
	if n.Comment != nil {
		Walk(v, n.Comment)
	}*/

	case *ast.FieldList:
	/*for _, f := range n.List {
		Walk(v, f)
	}*/

	// Expressions
	case *ast.BadExpr, *ast.Ident, *ast.BasicLit:
	// nothing to do

	case *ast.Ellipsis:
	/*if n.Elt != nil {
		Walk(v, n.Elt)
	}*/

	case *ast.FuncLit:
	/*Walk(v, n.Type)
	Walk(v, n.Body)*/

	case *ast.CompositeLit:
	/*if n.Type != nil {
		Walk(v, n.Type)
	}
	walkExprList(v, n.Elts)*/

	case *ast.ParenExpr:
	/*Walk(v, n.X)*/

	case *ast.SelectorExpr:
	/*Walk(v, n.X)
	Walk(v, n.Sel)*/

	case *ast.IndexExpr:
	/*Walk(v, n.X)
	Walk(v, n.Index)*/

	case *ast.SliceExpr:
	/*Walk(v, n.X)
	if n.Low != nil {
		Walk(v, n.Low)
	}
	if n.High != nil {
		Walk(v, n.High)
	}
	if n.Max != nil {
		Walk(v, n.Max)
	}*/

	case *ast.TypeAssertExpr:
	/*Walk(v, n.X)
	if n.Type != nil {
		Walk(v, n.Type)
	}*/

	case *ast.CallExpr:
	/*Walk(v, n.Fun)
	walkExprList(v, n.Args)*/

	case *ast.StarExpr:
	/*Walk(v, n.X)*/

	case *ast.UnaryExpr:
	/*Walk(v, n.X)*/

	case *ast.BinaryExpr:
	/*Walk(v, n.X)
	Walk(v, n.Y)*/

	case *ast.KeyValueExpr:
	/*Walk(v, n.Key)
	Walk(v, n.Value)*/

	// Types
	case *ast.ArrayType:
	/*if n.Len != nil {
		Walk(v, n.Len)
	}
	Walk(v, n.Elt)*/

	case *ast.StructType:
	/*Walk(v, n.Fields)*/

	case *ast.FuncType:
	/*if n.Params != nil {
		Walk(v, n.Params)
	}
	if n.Results != nil {
		Walk(v, n.Results)
	}*/

	case *ast.InterfaceType:
	/*Walk(v, n.Methods)*/

	case *ast.MapType:
	/*Walk(v, n.Key)
	Walk(v, n.Value)*/

	case *ast.ChanType:
	/*Walk(v, n.Value)*/

	// Statements
	case *ast.BadStmt:
	// nothing to do

	case *ast.DeclStmt:
	/*Walk(v, n.Decl)*/

	case *ast.EmptyStmt:
	// nothing to do

	case *ast.LabeledStmt:
	/*Walk(v, n.Label)
	Walk(v, n.Stmt)*/

	case *ast.ExprStmt:
	/*Walk(v, n.X)*/

	case *ast.SendStmt:
	/*Walk(v, n.Chan)
	Walk(v, n.Value)*/

	case *ast.IncDecStmt:
	/*Walk(v, n.X)*/

	case *ast.AssignStmt:
	/*walkExprList(v, n.Lhs)
	walkExprList(v, n.Rhs)*/

	case *ast.GoStmt:
	/*Walk(v, n.Call)*/

	case *ast.DeferStmt:
	/*Walk(v, n.Call)*/

	case *ast.ReturnStmt:
		/*walkExprList(v, n.Results)*/

	case *ast.BranchStmt:
	/*if n.Label != nil {
		Walk(v, n.Label)
	}*/

	case *ast.BlockStmt:
	/*walkStmtList(v, n.List)*/

	case *ast.IfStmt:
	/*if n.Init != nil {
		Walk(v, n.Init)
	}
	Walk(v, n.Cond)
	Walk(v, n.Body)
	if n.Else != nil {
		Walk(v, n.Else)
	}*/

	case *ast.CaseClause:
	/*walkExprList(v, n.List)
	walkStmtList(v, n.Body)*/

	case *ast.SwitchStmt:
	/*if n.Init != nil {
		Walk(v, n.Init)
	}
	if n.Tag != nil {
		Walk(v, n.Tag)
	}
	Walk(v, n.Body)*/

	case *ast.TypeSwitchStmt:
	/*if n.Init != nil {
		Walk(v, n.Init)
	}
	Walk(v, n.Assign)
	Walk(v, n.Body)*/

	case *ast.CommClause:
	/*if n.Comm != nil {
		Walk(v, n.Comm)
	}
	walkStmtList(v, n.Body)*/

	case *ast.SelectStmt:
	/*Walk(v, n.Body)*/

	case *ast.ForStmt:
	/*if n.Init != nil {
		Walk(v, n.Init)
	}
	if n.Cond != nil {
		Walk(v, n.Cond)
	}
	if n.Post != nil {
		Walk(v, n.Post)
	}
	Walk(v, n.Body)*/

	case *ast.RangeStmt:
	/*if n.Key != nil {
		Walk(v, n.Key)
	}
	if n.Value != nil {
		Walk(v, n.Value)
	}
	Walk(v, n.X)
	Walk(v, n.Body)*/

	// Declarations
	case *ast.ImportSpec:
	/*if n.Doc != nil {
		Walk(v, n.Doc)
	}
	if n.Name != nil {
		Walk(v, n.Name)
	}
	Walk(v, n.Path)
	if n.Comment != nil {
		Walk(v, n.Comment)
	}*/

	case *ast.ValueSpec:
	/*if n.Doc != nil {
		Walk(v, n.Doc)
	}
	walkIdentList(v, n.Names)
	if n.Type != nil {
		Walk(v, n.Type)
	}
	walkExprList(v, n.Values)
	if n.Comment != nil {
		Walk(v, n.Comment)
	}*/

	case *ast.TypeSpec:
	/*if n.Doc != nil {
		Walk(v, n.Doc)
	}
	Walk(v, n.Name)
	Walk(v, n.Type)
	if n.Comment != nil {
		Walk(v, n.Comment)
	}*/

	case *ast.BadDecl:
	// nothing to do

	case *ast.GenDecl:
	/*if n.Doc != nil {
		Walk(v, n.Doc)
	}
	for _, s := range n.Specs {
		Walk(v, s)
	}*/

	case *ast.FuncDecl:
	/*if n.Doc != nil {
		Walk(v, n.Doc)
	}
	if n.Recv != nil {
		Walk(v, n.Recv)
	}
	Walk(v, n.Name)
	Walk(v, n.Type)
	if n.Body != nil {
		Walk(v, n.Body)
	}*/

	// Files and packages
	case *ast.File:
	/*if n.Doc != nil {
		Walk(v, n.Doc)
	}
	Walk(v, n.Name)
	walkDeclList(v, n.Decls)*/
	// don't walk n.Comments - they have been
	// visited already through the individual
	// nodes

	case *ast.Package:
	/*for _, f := range n.Files {
		Walk(v, f)
	}*/

	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}
	return true
}
