package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path"
	"path/filepath"
	"strings"
)

// 所有GO文件基础信息
var infoList []fileInfo

type fileInfo struct {
	fileBaseName string //文件名
	fileRelName  string //相对路径名
	fileAbsName  string //绝对路径名
	PkgName      string //ny
	imports      []string
	cons         []string
	structName   []string
	funcName     []string //普通函数名
	method       []method
}

type method struct {
	PkgName    string //ny
	StructName string //ny
	Receiver   string
	methodName string
}

// 处理路径下所有文件基础信息
func findFileInfos() {
	for _, filename := range files {
		findFileInfo(filename)
	}
}

// 处理单个文件的基础信息
func findFileInfo(filename string) {

	info := fileInfo{}

	//处理文件名
	baseName := filepath.Base(filename)
	info.fileBaseName = baseName
	relName, _ := filepath.Rel(root, filename)
	info.fileRelName = relName
	info.fileAbsName = filename

	f, _ := astParser(filename)
	//使用类型断言时，将*ast.File放置在一个接口类型的变量中
	var i interface{} = f
	switch n := i.(type) {
	case *ast.File:
		for _, d := range n.Decls {
			switch decl := d.(type) {
			case *ast.GenDecl:
				switch decl.Tok {
				case token.CONST:
					for _, spec := range decl.Specs {
						vspec := spec.(*ast.ValueSpec)
						for _, id := range vspec.Names {
							info.cons = append(info.cons, id.Name)
						}
					}
				case token.TYPE:
					for _, spec := range decl.Specs {
						tspec := spec.(*ast.TypeSpec)
						info.structName = append(info.structName, tspec.Name.Name)
					}
				}
			case *ast.FuncDecl:
				info.funcName = append(info.funcName, decl.Name.Name)
				if decl.Recv != nil {
					for _, field := range decl.Recv.List {
						if t, ok := field.Type.(*ast.StarExpr); ok {
							if ident, ok := t.X.(*ast.Ident); ok {
								methodTemp := method{
									Receiver:   ident.Name,
									methodName: decl.Name.Name,
								}
								info.method = append(info.method, methodTemp)
							}
						}
					}
				}
			}
		}
		//处理imports
		for _, imp := range f.Imports {
			impPath := strings.Trim(imp.Path.Value, "\"") // 删除引号
			impName := path.Base(impPath)                 // 提取包名
			info.imports = append(info.imports, impName)
		}
	}
	infoList = append(infoList, info)
}

func printFileInfo(info fileInfo) {
	fmt.Printf("FileBaseName: %s\n", info.fileBaseName)
	fmt.Printf("Import: %s\n", info.imports)
	fmt.Printf("Constants: %v\n", info.cons)
	fmt.Printf("Structs: %v\n", info.structName)
	fmt.Printf("Functions: %v\n", info.funcName)
	for _, m := range info.method {
		fmt.Printf("Methods: %s.%s\n", m.Receiver, m.methodName)
	}
}

// 解析指定的Go代码文件，返回一个ast.File类型的对象 表示整个源代码文件的抽象语法树
func astParser(filename string) (*ast.File, *token.FileSet) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	return f, fset
}
