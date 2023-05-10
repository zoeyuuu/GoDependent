package tool

import "go/ast"

func findFunctionCall(n *ast.CallExpr, v *Visitor) {
	switch f := n.Fun.(type) {
	case *ast.Ident:
		// 处理函数调用
		pos := v.fset.Position(n.Pos())
		name := n.Fun.(*ast.Ident).Name
		// 包名不同必不是调用
		if infoList[v.I].PkgName != infoList[v.J].PkgName {
			return
		}
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
		name := n.Fun.(*ast.Ident).Name
		switch f.X.(type) {
		case *ast.Ident:
			// 两个甚至三个的时候 通过包名比较来判断 还是通过ident.obj?=nil
			if infoList[v.J].PkgName == f.X.(*ast.Ident).Name {

			}
		case *ast.SelectorExpr:
			// 处理嵌套的 selectorexpr
			// ...
		}
	}
}

/*
case *ast.SelectorExpr:
    // 处理嵌套调用
    name := n.Fun.(*ast.Ident).Name
    switch f.X.(type) {
    case *ast.Ident:
        // 左侧是一个标识符
        ident := f.X.(*ast.Ident)
        if infoList[v.J].PkgName == ident.Name {
            // 在同一个包中
        } else {
            // 不在同一个包中，需要进一步判断
            if ident.Obj == nil {
                // 左侧是一个未声明的标识符，可能是外部包中的函数
                // ...
            } else {
                // 左侧是一个已声明的标识符，需要递归判断
                findFunctionCall(&ast.CallExpr{Fun: ident}, v)
            }
        }
    case *ast.SelectorExpr:
        // 左侧是一个嵌套的 SelectorExpr，继续递归处理
        // ...
    }
}
*/
