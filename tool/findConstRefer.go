package tool

import (
	"go/ast"
)

// 还是有问题 所有的ident会有不同包重名问题
func findConstRefer(n *ast.Ident, v *Visitor) {
	if n.Obj != nil {
		// obj.kind有四种
		if n.Obj.Kind == ast.Con {
			constName := n.Obj.Name
			for _, vr := range infoList[v.K].Cons {
				if constName == vr.Name {
					temp := constRefer{
						name: constName,
					}
					v.Dep.Relations["ConstRefer"] = append(v.Dep.Relations["ConstRefer"], temp)
				}
			}
		}
		if n.Obj.Kind == ast.Var {
			varName := n.Obj.Name
			if varName != "_" {
				for i, vr := range infoList[v.K].Vars {
					if varName == vr.Name {
						temp := globalRefer{
							name: varName,
							Type: infoList[v.K].Vars[i].Kind,
						}
						v.Dep.Relations["GlobalRefer"] = append(v.Dep.Relations["GlobalRefer"], temp)
					}
				}
			}
		}
	}
}
