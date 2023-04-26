package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type selector struct {
	left  string    // 选择器表达式所属的变量的名称 a
	right string    // 选择器选择的字段或方法名称 b
	pos   token.Pos // 源代码中的位置pos
	line  string    // 所在行line
}

func (s *selector) toString() string {
	return s.left + "." + s.right
}

type allSelectorVisitor struct {
	selectors []selector
}

func main() {
	filename := "D:\\projects\\go_projects\\src\\GoDependent\\file\\b.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	varAll := allSelectorVisitor{}
	ast.Walk(&varAll, f)
	ast.Walk(&visitor{}, f)
	fmt.Println(varAll)
}

// visitor 实现 ast.Visitor 接口
type visitor struct {
	k int
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	// 处理语法树节点
	switch n := node.(type) {
	case *ast.GenDecl:
		if n.Tok == token.VAR {
			for _, spec := range n.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					// typename对应的结构体名
					ident := valueSpec.Type.(*ast.Ident)
					fmt.Println(ident)
				}
			}
		}

	case *ast.CallExpr:
		fmt.Printf("found call to %s\n", n.Fun)
	}
	return v
}

func (v *allSelectorVisitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return v
	}
	// 判断它是否是一个 SelectorExpr 类型的表达式（即变量选择器）
	if selectorExp, ok := n.(*ast.SelectorExpr); ok {
		// 只判断a.b
		fmt.Println(selectorExp.Sel)
		if va, ok := selectorExp.X.(*ast.Ident); ok {
			if va.Obj == nil {
				return v
			}
			if va.Obj.Kind.String() == "var" {

				newSelector := selector{
					left:  va.Name,              //选择器表达式所属的变量的名称
					right: selectorExp.Sel.Name, // 选择器选择的字段或方法名称
					pos:   va.Pos(),
				}
				v.add(newSelector)
			}
		}
	}
	return v
}

func (v *allSelectorVisitor) add(s selector) {
	if !v.exists(s) {
		v.selectors = append(v.selectors, s)
	}
}

func (v *allSelectorVisitor) exists(s selector) bool {
	for _, n := range v.selectors {
		if n.left == s.left && n.right == s.right {
			return true
		}
	}
	return false
}
