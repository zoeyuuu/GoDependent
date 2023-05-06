package tool

import (
	"GoDependent/file"
	"fmt"
	"go/ast"
)

var infoList []file.FileInfo

type Dependencies struct {
	Src       string
	Des       string
	Relations map[string][]any
}

var DependencyList []Dependencies

func FindDependenyAll() {
	infoList = file.InfoList
	for i := 0; i < len(infoList); i++ {
		for j := 0; j < len(infoList); j++ {
			if i != j {
				findDependency(i, j)
			}
		}
	}
}

func findDependency(i, j int) {
	filename := infoList[i].FileAbsName
	f, _ := file.AstParser(filename)
	// 两个文件间单向的依赖关系
	dependency := &Dependencies{
		Src:       infoList[i].FileRelName,
		Des:       infoList[j].FileRelName,
		Relations: make(map[string][]any),
	}
	v := &Visitor{K: j, Dep: dependency}
	ast.Walk(v, f)
	if len(dependency.Relations) != 0 {
		DependencyList = append(DependencyList, *dependency)
	}
	if _, ok := dependency.Relations["GlobalRefer"]; ok {
		fmt.Println(dependency)
	}
}

func FindDependencytest(filename1, filename2 string) {
	var i, j int
	infoList = file.InfoList
	for k, v := range infoList {
		if v.FileBaseName == filename1 {
			i = k
		}
		if v.FileBaseName == filename2 {
			j = k
		}
	}
	filename := infoList[i].FileAbsName
	f, fset := file.AstParser(filename)
	// 两个文件间单向的依赖关系
	dependency := &Dependencies{
		Src:       infoList[i].FileRelName,
		Des:       infoList[j].FileRelName,
		Relations: make(map[string][]any),
	}
	v := &Visitor{K: j, Dep: dependency, fset: fset}
	ast.Walk(v, f)
	if len(dependency.Relations) != 0 {
		DependencyList = append(DependencyList, *dependency)
	}
}
