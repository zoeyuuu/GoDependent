package main

import (
	"go/ast"
	"go/token"
)

// 查找var语句实例化
func findVar(n ast.Node) bool {
	// 将参数 n 转换为 *ast.GenDecl 类型
	genDecl, ok := n.(*ast.GenDecl)
	if !ok || genDecl.Tok != token.VAR {
		return true
	}
	for _, spec := range genDecl.Specs {
		if valueSpec, ok := spec.(*ast.ValueSpec); ok {
			// typename对应的结构体名
			typename := valueSpec.Type.(*ast.Ident).Name
			for _, v := range valueSpec.Names {
				// 存储多个实例化变量名
				mp[v.Name] = typename
				// for _, v := range info2.structName
				for _, v := range infoList[0].structName {
					if typename == v {
						dependency.relation["instantiation"]++
					}
				}
			}
		}
	}
	return true
}

// 查找结构体继承
func findInheritance(n ast.Node) bool {
	typeSpec, ok := n.(*ast.TypeSpec)
	if !ok {
		return true
	}
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return true
	}
	for _, field := range structType.Fields.List {
		ident, ok := field.Type.(*ast.Ident)
		if !ok {
			continue
		}
		for _, v := range infoList[1].structName {
			if ident.Name == v {
				dependency.relation["Inheritance"]++
			}
		}
	}
	return true
}

// 查找普通函数调用和方法调用
func findFunction(n ast.Node) bool {
	// 将参数 n 转换为 **ast.ExprStmt 类型
	exprStmt, ok := n.(*ast.ExprStmt)
	if !ok {
		return true
	}
	// 检查 ExprStmt 的表达式是否是 CallExpr
	if call, ok := exprStmt.X.(*ast.CallExpr); ok {

		// **检查是否是调用结构体方法
		if ident, ok := call.Fun.(*ast.SelectorExpr); ok {
			if _, ok := mp[ident.X.(*ast.Ident).Name]; ok {
				// _是对于结构体名
				dependency.relation["call_method"]++
			}
		}

		// **检查是否是普通函数
		if ident, ok := call.Fun.(*ast.Ident); ok {
			for _, v := range infoList[1].funcName {
				if v == ident.Name {
					dependency.relation["call_normal"]++
				}
			}
		}
	}
	return true
}
