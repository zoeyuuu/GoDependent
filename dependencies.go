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
	info1 := infoList[0] //b.go
	info2 := infoList[1] //a.go
	dependency.src = info1.fileName
	dependency.des = info2.fileName
	dependency.relation = make(map[string]int)
	findRelationshipsAB()
	//return denpendency
}

func findRelationshipsAB() {
	filename1 := "./file/b.go"
	f, _ := astParser(filename1)

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
