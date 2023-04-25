package main

import (
	"fmt"
	"go/ast"
)

type dependencies struct {
	src      string
	des      string
	relation map[string]int
}

var mp map[string]string

var dependency dependencies

func findDependencies() {
	for i := 0; i < len(infoList); i++ {
		for j := 0; j < len(infoList); j++ {
			if i != j {
				findRelationshipsAtoB(i, j)
			}
		}
	}
	info1 := infoList[0] //b.go
	info2 := infoList[1] //a.go
	dependency.src = info1.fileRelName
	dependency.des = info2.fileRelName
	dependency.relation = make(map[string]int)

}

func findRelationshipsAtoB(i, j int) {
	filename := infoList[i].fileAbsName
	f, _ := astParser(filename)

	// 初始化结构体-示例化map
	mp = make(map[string]string)

	// astinspect
	// 查找实例化
	ast.Inspect(f, findVar)
	ast.Inspect(f, findInheritance)
	ast.Inspect(f, findFunction)

}

func printDenpendencies(dep dependencies) {
	fmt.Println(dep.src, "->", dep.des)
	for k, v := range dep.relation {
		fmt.Println(k, ":", v)
	}
}
