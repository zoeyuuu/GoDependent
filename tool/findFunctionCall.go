package tool

import "go/ast"

func findFunctionCall(n *ast.CallExpr, v *Visitor) {
	samePackage := infoList[v.I].PkgName == infoList[v.J].PkgName
	pos := v.fset.Position(n.Pos())
	switch f := n.Fun.(type) {
	case *ast.Ident:
		// 包名不同必不是调用
		if !samePackage {
			return
		}
		name := n.Fun.(*ast.Ident).Name
		for _, funName := range infoList[v.J].FuncName {
			if name == funName {
				tmp := functionCall{
					funName: name,
					pos:     pos,
				}
				v.Dep.Relations["FunctionCall"] = append(v.Dep.Relations["FunctionCall"], tmp)
			}
		}
	case *ast.SelectorExpr:
		// 处理嵌套调用
		name := n.Fun.(*ast.SelectorExpr).Sel.Name
		switch f.X.(type) {
		// A.B()情况
		case *ast.Ident:
			//不同包情况：A是其他包名
			if !samePackage {
				if infoList[v.J].PkgName == f.X.(*ast.Ident).Name {
					for _, funName := range infoList[v.J].FuncName {
						if name == funName {
							tmp := functionCall{
								funName: name,
								pos:     pos,
							}
							v.Dep.Relations["FunctionCall"] = append(v.Dep.Relations["FunctionCall"], tmp)
						}
					}
				}
			} else { //同一个包：A是实例化结构体名 只需比较Fun.Sel.Name
				for _, funName := range infoList[v.J].FuncName {
					if name == funName {
						tmp := functionCall{
							funName: name,
							pos:     pos,
						}
						v.Dep.Relations["FunctionCall"] = append(v.Dep.Relations["FunctionCall"], tmp)
					}
				}
			}
		//A.B.C()
		case *ast.SelectorExpr:
			// 报错f.X.(*ast.SelectorExpr).X可能是*ast.CallExpr类型？
			if _, ok := f.X.(*ast.SelectorExpr).X.(*ast.Ident); ok {
				if infoList[v.J].PkgName == f.X.(*ast.SelectorExpr).X.(*ast.Ident).Name {
					for _, funName := range infoList[v.J].FuncName {
						if name == funName {
							tmp := functionCall{
								funName: name,
								pos:     pos,
							}
							v.Dep.Relations["FunctionCall"] = append(v.Dep.Relations["FunctionCall"], tmp)
						}
					}
				}
			}
			//
		}
	}
}
