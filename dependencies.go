package main

import (
	"fmt"
	"go/ast"
)

type dependencies struct {
	src       string
	des       string
	relation  map[string]int
	relations map[string]struct{}
}

var dependencyList []dependencies

func findDependencies() {

	for i := 0; i < len(infoList); i++ {
		for j := 0; j < len(infoList); j++ {
			if i != j {
				findDependencyAtoB(i, j)
			}
		}
	}
	//dependency.relation = make(map[string]int)

}

func findDependencyAtoB(i, j int) {
	var dependency dependencies
	dependency.src = infoList[i].fileRelName
	dependency.des = infoList[j].fileRelName

	filename := infoList[i].fileAbsName
	f, _ := astParser(filename)

	// 存储实例化
	var mp map[string]string
	// 初始化结构体-示例化map
	mp = make(map[string]string)

	// astinspect
	// 查找实例化
	ast.Inspect(f, findVar)
	ast.Inspect(f, findInheritance)
	ast.Inspect(f, findFunction)

	dependencyList = append(dependencyList, dependency)

}

func printDenpendenyList() {
	for _, dep := range dependencyList {
		fmt.Println(dep.src, "->", dep.des)
		for k, v := range dep.relation {
			fmt.Println(k, ":", v)
		}
	}
}
