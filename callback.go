package main

import (
	"go/ast"
	"go/token"
)

// 查找var语句实例化
func findVar(n *ast.GenDecl, v *visitor) {
	if n.Tok == token.VAR {
		for _, spec := range n.Specs {
			if valueSpec, ok := spec.(*ast.ValueSpec); ok {
				switch node := valueSpec.Type.(type) {
				case *ast.Ident:
					// 普通类型
					for _, structName := range infoList[v.k].structName {
						//类型匹配成功
						if node.Name == structName {
							for _, identName := range valueSpec.Names {
								inst := instantiation{typeName: structName, varName: identName.Name}
								v.dep.relations["instantiation"] = append(v.dep.relations["instantiation"], inst)
							}
						}
					}
				case *ast.SelectorExpr:
					// 跨包类型 sel一定是Ident类型 暂时认定X也是
					if infoList[v.k].PkgName == node.X.(*ast.Ident).Name {
						for _, structName := range infoList[v.k].structName {
							//类型匹配成功
							if structName == node.Sel.Name {
								for _, identName := range valueSpec.Names {
									inst := instantiation{typeName: structName, varName: identName.Name}
									v.dep.relations["instantiation"] = append(v.dep.relations["instantiation"], inst)
								}
							}
						}
					}
				}
			}
		}
	}
}

// 查找继承(嵌套匿名结构体)
func findStructRelation(n *ast.GenDecl, v *visitor) {
	// 处理结构体类型
	if n.Tok == token.TYPE {
		for _, spec := range n.Specs {
			if typeSpec, ok := spec.(*ast.TypeSpec); ok {
				if structType, ok := typeSpec.Type.(*ast.StructType); ok {
					// field是struct所有声明
					for _, field := range structType.Fields.List {
						switch node := field.Type.(type) {
						case *ast.Ident:
							for _, structName := range infoList[v.k].structName {
								// 匹配成功
								if node.Name == structName {
									//匿名嵌套
									if field.Names == nil {
										temp := structEmbedding{
											container: typeSpec.Name.Name,
											member:    field.Type.(*ast.Ident).Name,
										}
										v.dep.relations["structEmbedding"] = append(v.dep.relations["structEmbedding"], temp)
									} else {
										//聚合关系
										for _, identName := range field.Names {
											temp := structAggregation{
												whole:   typeSpec.Name.Name,
												part:    field.Type.(*ast.Ident).Name,
												varName: identName.Name,
											}
											v.dep.relations["structAggregation"] = append(v.dep.relations["structAggregation"], temp)
										}
									}
								}
							}
						case *ast.SelectorExpr:
							// 包名匹配
							if infoList[v.k].PkgName == node.X.(*ast.Ident).Name {
								for _, structName := range infoList[v.k].structName {
									// 匹配成功
									if node.Sel.Name == structName {
										//匿名嵌套
										if field.Names == nil {
											temp := structEmbedding{
												container: typeSpec.Name.Name,
												member:    field.Type.(*ast.SelectorExpr).Sel.Name,
											}
											v.dep.relations["structEmbedding"] = append(v.dep.relations["structEmbedding"], temp)
										} else {
											//聚合关系
											for _, identName := range field.Names {
												temp := structAggregation{
													whole:   typeSpec.Name.Name,
													part:    field.Type.(*ast.SelectorExpr).Sel.Name,
													varName: identName.Name,
												}
												v.dep.relations["structAggregation"] = append(v.dep.relations["structAggregation"], temp)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

/*
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
*/
