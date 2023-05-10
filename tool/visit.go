package tool

import (
	"go/ast"
	"go/token"
)

type Visitor struct {
	I        int
	J        int // j作为visitor结构体的字段
	Dep      *Dependencies
	comments ast.CommentMap
	fset     *token.FileSet
	f        *ast.File
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {

	// 处理语法树节点
	switch n := node.(type) {
	case *ast.GenDecl:
		//var实例化
		findVarInst(n, v)
		//结构体嵌套、组合
		findStructRelation(n, v)
		//接口嵌套
		//findInterfaceRelation(n,v)
	case *ast.CallExpr:
		//查找函数/方法调用
		findFunctionCall(n, v)
	case *ast.Ident:
		// 排除注释里的情况
		if _, ok := v.comments[&ast.Ident{NamePos: n.Pos()}]; ok {
			return v
		}
		// 找到
		findConstRefer(n, v)
	}
	return v
}

/*
func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	// 获取当前节点的位置信息
	pos := v.fset.Position(node.Pos())
	// 获取该位置的注释信息
	commentGroup := v.Info.Comments[pos]
	// 如果有注释，并且不是行注释或块注释，则进行处理
	if commentGroup != nil && !strings.HasPrefix(commentGroup.Text(), "//") && !strings.HasPrefix(commentGroup.Text(), "/*") {
		switch n := node.(type) {
		case *ast.GenDecl:
			//var实例化
			findVar(n, v)
			//结构体嵌套、组合
			findStructRelation(n, v)
			//接口嵌套
			//findInterfaceRelation(n,v)
		case *ast.Ident:
			findConstRefer(n, v)
		}
	}
	return v
}
*/
