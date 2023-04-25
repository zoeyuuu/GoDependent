package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
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
	fset := token.NewFileSet()
	// 解析指定的Go代码文件，返回一个ast.File类型的对象 表示整个源代码文件的抽象语法树
	f, err := parser.ParseFile(fset, filename1, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
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
