package main

import "go/ast"

type visitor struct {
	k   int // j作为visitor结构体的字段
	dep *Dependencies
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	// 处理语法树节点
	switch n := node.(type) {
	case *ast.GenDecl:
		findVar(n, v)
		findStructRelation(n, v)
	}
	return v
}
