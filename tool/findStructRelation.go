package tool

import (
	"go/ast"
	"go/token"
)

// 查找继承(嵌套匿名结构体)
func findStructRelation(n *ast.GenDecl, v *Visitor) {
	// 处理结构体类型
	if n.Tok == token.TYPE {
		for _, spec := range n.Specs {
			if typeSpec, ok := spec.(*ast.TypeSpec); ok {
				if structType, ok := typeSpec.Type.(*ast.StructType); ok {
					// field是struct所有声明
					for _, field := range structType.Fields.List {
						switch node := field.Type.(type) {
						case *ast.Ident:
							for _, structName := range infoList[v.J].StructName {
								// 匹配成功
								if node.Name == structName {
									//匿名嵌套
									if field.Names == nil {
										pos := v.fset.Position(node.Pos())
										temp := structEmbedding{
											Container: typeSpec.Name.Name,
											Member:    field.Type.(*ast.Ident).Name,
											Pos:       pos,
										}
										v.Dep.Relations["structEmbedding"] = append(v.Dep.Relations["structEmbedding"], temp)
									} else {
										//聚合关系
										for _, identName := range field.Names {
											pos := v.fset.Position(node.Pos())
											temp := structAggregation{
												Whole:   typeSpec.Name.Name,
												Part:    field.Type.(*ast.Ident).Name,
												ObjName: identName.Name,
												Pos:     pos,
											}
											v.Dep.Relations["structAggregation"] = append(v.Dep.Relations["structAggregation"], temp)
										}
									}
								}
							}
						case *ast.SelectorExpr:
							// 包名匹配
							if infoList[v.J].PkgName == node.X.(*ast.Ident).Name {
								for _, structName := range infoList[v.J].StructName {
									// 匹配成功
									if node.Sel.Name == structName {
										//匿名嵌套
										if field.Names == nil {
											pos := v.fset.Position(node.Pos())
											temp := structEmbedding{
												Container: typeSpec.Name.Name,
												Member:    field.Type.(*ast.SelectorExpr).Sel.Name,
												Pos:       pos,
											}
											v.Dep.Relations["structEmbedding"] = append(v.Dep.Relations["structEmbedding"], temp)
										} else {
											//聚合关系
											for _, identName := range field.Names {
												pos := v.fset.Position(node.Pos())
												temp := structAggregation{
													Whole:   typeSpec.Name.Name,
													Part:    field.Type.(*ast.SelectorExpr).Sel.Name,
													ObjName: identName.Name,
													Pos:     pos,
												}
												v.Dep.Relations["structAggregation"] = append(v.Dep.Relations["structAggregation"], temp)
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
