package tool

import (
	"go/ast"
	"go/token"
)

// 查找var语句实例化
func findVar(n *ast.GenDecl, v *Visitor) {
	if n.Tok == token.VAR {
		for _, spec := range n.Specs {
			if valueSpec, ok := spec.(*ast.ValueSpec); ok {
				switch node := valueSpec.Type.(type) {
				case *ast.Ident:
					// 普通类型
					for _, structName := range infoList[v.K].StructName {
						//类型匹配成功
						if node.Name == structName {
							for _, identName := range valueSpec.Names {
								inst := Instantiation{TypeName: structName, VarName: identName.Name}
								v.Dep.Relations["instantiation"] = append(v.Dep.Relations["instantiation"], inst)
							}
						}
					}
				case *ast.SelectorExpr:
					// 跨包类型 sel一定是Ident类型 暂时认定X也是
					if infoList[v.K].PkgName == node.X.(*ast.Ident).Name {
						for _, structName := range infoList[v.K].StructName {
							//类型匹配成功
							if structName == node.Sel.Name {
								for _, identName := range valueSpec.Names {
									inst := Instantiation{TypeName: structName, VarName: identName.Name}
									v.Dep.Relations["instantiation"] = append(v.Dep.Relations["instantiation"], inst)
								}
							}
						}
					}
				}
			}
		}
	}
}
