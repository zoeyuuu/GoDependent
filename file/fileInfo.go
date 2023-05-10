package file

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path"
	"path/filepath"
	"strings"
)

// 所有GO文件基础信息
var InfoList []FileInfo

// 思考一下
// type structName string

type FileInfo struct {
	FileBaseName  string //文件名
	FileRelName   string //相对路径名
	FileAbsName   string //绝对路径名
	PkgName       string //ny
	imports       []string
	Cons          []Con
	structTypes   []StructType
	funcTypes     []FuncType
	interfaceType []InterfaceType
	arrayTypes    []ArrayType
	Vars          []Var
	//	Type         []any
	StructName []string
	FuncName   []string //普通函数名
	method     []method
}

type Con struct {
	Name string
	Kind string
}

type Var struct {
	Name string
	Kind string
}
type StructType struct {
	name string
}
type FuncType struct {
	name string
}
type InterfaceType struct {
	// 继承要不要考虑
	name        string
	methodNames []string
	// methods []method
}
type ArrayType struct {
	name string
}

type method struct {
	PkgName    string //ny
	StructName string //ny
	Receiver   string
	methodName string
}

// 处理路径下所有文件基础信息
func FindFileInfos() {
	for _, filename := range Files {
		findFileInfo(filename)
	}
}

// 处理单个文件的基础信息
func findFileInfo(filename string) {

	info := FileInfo{}

	//处理文件名
	baseName := filepath.Base(filename)
	info.FileBaseName = baseName
	relName, _ := filepath.Rel(root, filename)
	// 反斜杠转换为斜杠
	relName = filepath.ToSlash(relName)
	info.FileRelName = relName
	info.FileAbsName = filename

	f, _ := AstParser(filename)
	// 包名
	info.PkgName = f.Name.Name

	//使用类型断言时，将*ast.File放置在一个接口类型的变量中
	var i interface{} = f
	switch n := i.(type) {
	case *ast.File:
		for _, d := range n.Decls {
			switch decl := d.(type) {
			// GenDecl.TOKEN:IMPORT, CONST, TYPE, or VAR
			case *ast.GenDecl:
				switch decl.Tok {
				case token.CONST:
					for _, spec := range decl.Specs {
						valueSpec := spec.(*ast.ValueSpec)
						for i, name := range valueSpec.Names {
							con := Con{Name: name.Name}
							if valueSpec.Values != nil {
								if basicLit, ok := valueSpec.Values[i].(*ast.BasicLit); ok {
									con.Kind = basicLit.Kind.String()
								} else {
									// 处理不是 *ast.BasicLit 的情况
									con.Kind = "unknown"
								}
							}
							info.Cons = append(info.Cons, con)
						}
					}
				case token.VAR:
					for _, spec := range decl.Specs {
						vspec := spec.(*ast.ValueSpec)
						for _, id := range vspec.Names {
							tmp := Var{Name: id.Name}
							if vspec.Type != nil {
								if ident, ok := vspec.Type.(*ast.Ident); ok {
									tmp.Kind = ident.Name
								} else {
									// 处理不是 *ast.Ident 的情况
									tmp.Name = "unknown"
								}
							}
							info.Vars = append(info.Vars, tmp)
						}
					}
				case token.TYPE:
					for _, spec := range decl.Specs {
						tspec := spec.(*ast.TypeSpec)
						switch t := tspec.Type.(type) {
						case *ast.StructType:
							info.structTypes = append(info.structTypes, StructType{name: tspec.Name.Name})
						case *ast.FuncType:
							info.funcTypes = append(info.funcTypes, FuncType{name: tspec.Name.Name})
						case *ast.InterfaceType:
							inter := InterfaceType{name: tspec.Name.Name}
							//处理接口的内容信息 可以是方法/继承 继承先不考虑
							for _, field := range t.Methods.List {
								if _, ok := field.Type.(*ast.FuncType); ok {
									for _, identName := range field.Names {
										inter.methodNames = append(inter.methodNames, identName.Name)
									}
								}
							}
							info.interfaceType = append(info.interfaceType, inter)
						case *ast.ArrayType:
							info.arrayTypes = append(info.arrayTypes, ArrayType{name: tspec.Name.Name})
						}
					}
					for _, spec := range decl.Specs {
						tspec := spec.(*ast.TypeSpec)
						info.StructName = append(info.StructName, tspec.Name.Name)
					}
				}
			case *ast.FuncDecl:
				info.FuncName = append(info.FuncName, decl.Name.Name)
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

	InfoList = append(InfoList, info)
}

// 解析指定的Go代码文件，返回一个ast.File类型的对象 表示整个源代码文件的抽象语法树
func AstParser(filename string) (*ast.File, *token.FileSet) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	return f, fset
}
