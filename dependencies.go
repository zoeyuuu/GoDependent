package main

import (
	"go/ast"
)

type dependencies struct {
	src       string
	des       string
	relations map[string][]any
}

type visitor struct {
	k   int // j作为visitor结构体的字段
	dep *dependencies
}

var dependencyList []dependencies

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
	dependency := &dependencies{
		src:       infoList[i].fileRelName,
		des:       infoList[j].fileRelName,
		relations: make(map[string][]any),
	}
	v := &visitor{k: j, dep: dependency}
	ast.Walk(v, f)
	dependencyList = append(dependencyList, *dependency)
}
