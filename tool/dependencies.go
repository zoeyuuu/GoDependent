package tool

import (
	"GoDependent/file"
	"go/ast"
)

var infoList []file.FileInfo

type Dependencies struct {
	Src         string `json:"abs_source"`
	Tar         string `json:"abs_target"`
	RelativeSrc string `json:"source"`
	RelativeTar string `json:"target"`
	Relations   map[string][]interface{}
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
	f, fset := file.AstParser(filename)
	comments := ast.NewCommentMap(fset, f, f.Comments)
	// 两个文件间单向的依赖关系
	dependency := &Dependencies{
		Src:         infoList[i].FileAbsName,
		RelativeSrc: infoList[i].FileRelName,
		Tar:         infoList[j].FileAbsName,
		RelativeTar: infoList[j].FileRelName,
		Relations:   make(map[string][]interface{}),
	}
	v := &Visitor{
		J:        j,
		I:        i,
		Dep:      dependency,
		comments: comments,
		f:        f,
		fset:     fset,
	}
	ast.Walk(v, f)
	if len(dependency.Relations) != 0 {
		DependencyList = append(DependencyList, *dependency)
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
		Tar:       infoList[j].FileRelName,
		Relations: make(map[string][]interface{}),
	}
	v := &Visitor{J: j, Dep: dependency, fset: fset}
	ast.Walk(v, f)
	if len(dependency.Relations) != 0 {
		DependencyList = append(DependencyList, *dependency)
	}
}
