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
			for _, vr := range infoList[v.J].Cons {
				if constName == vr.Name {
					pos := v.fset.Position(n.Pos()) // 获取引用点的位置信息
					temp := constRefer{
						name: constName,
						pos:  pos,
					}
					v.Dep.Relations["ConstRefer"] = append(v.Dep.Relations["ConstRefer"], temp)
				}
			}
		}
		if n.Obj.Kind == ast.Var {
			varName := n.Obj.Name
			// 排除“_”
			if varName == "_" {
				return
			}
			// 排除本文件内有
			for _, vr := range infoList[v.I].Vars {
				if varName == vr.Name {
					return
				}
			}
			// 先处理包内
			if infoList[v.I].PkgName != infoList[v.J].PkgName {
				return
			}
			for i, vr := range infoList[v.J].Vars {
				if varName == vr.Name {
					pos := v.fset.Position(n.Pos()) // 获取引用点的位置信息
					temp := globalRefer{
						name: varName,
						Type: infoList[v.J].Vars[i].Kind,
						pos:  pos,
					}
					v.Dep.Relations["GlobalRefer"] = append(v.Dep.Relations["GlobalRefer"], temp)
				}
			}
		}
	}
}
