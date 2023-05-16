package tool

import (
	"go/ast"
	"go/token"
)

// 查找var语句实例化
func findVarInst(n *ast.GenDecl, v *Visitor) {
	if n.Tok == token.VAR {
		for _, spec := range n.Specs {
			if valueSpec, ok := spec.(*ast.ValueSpec); ok {
				switch node := valueSpec.Type.(type) {
				case *ast.Ident:
					// 普通类型
					for _, structName := range infoList[v.J].StructName {
						//类型匹配成功
						if node.Name == structName {
							for _, identName := range valueSpec.Names {
								pos := v.fset.Position(n.Pos())
								tmp := Instantiation{
									StructName: structName,
									ObjName:    identName.Name,
									Token:      "var",
									Pos:        pos,
								}
								v.Dep.Relations["instantiation"] = append(v.Dep.Relations["instantiation"], tmp)
							}
						}
					}
				case *ast.SelectorExpr:
					// 跨包类型 sel一定是Ident类型 暂时认定X也是
					if infoList[v.J].PkgName == node.X.(*ast.Ident).Name {
						for _, structName := range infoList[v.J].StructName {
							//类型匹配成功
							if structName == node.Sel.Name {
								for _, identName := range valueSpec.Names {
									pos := v.fset.Position(n.Pos())
									tmp := Instantiation{
										StructName: structName,
										ObjName:    identName.Name,
										Token:      "var",
										Pos:        pos,
									}
									v.Dep.Relations["instantiation"] = append(v.Dep.Relations["instantiation"], tmp)
								}
							}
						}
					}
				}
			}
		}
	}
}
