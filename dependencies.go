package main

import (
	"go/ast"
)

type Dependencies struct {
	src       string
	des       string
	relations map[string][]any
}

var dependencyList []Dependencies

func findDependenyAll() {
	for i := 0; i < len(infoList); i++ {
		for j := 0; j < len(infoList); j++ {
			if i != j {
				findDependency(i, j)
			}
		}
	}
}

func findDependency(i, j int) {

	filename := infoList[i].fileAbsName
	f, _ := astParser(filename)
	// 两个文件间单向的依赖关系
	dependency := &Dependencies{
		src:       infoList[i].fileRelName,
		des:       infoList[j].fileRelName,
		relations: make(map[string][]any),
	}
	v := &visitor{k: j, dep: dependency}
	ast.Walk(v, f)
	if len(dependency.relations) != 0 {
		dependencyList = append(dependencyList, *dependency)
	}
}
