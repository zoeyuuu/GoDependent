package tool

import "go/token"

// 实例化
type Instantiation struct {
	StructName string
	ObjName    string
	Token      string // var or :=
	Pos        token.Position
}

// 结构体聚合
type structAggregation struct {
	Whole   string
	Part    string
	ObjName string
	Pos     token.Position
}

// 结构体嵌套
type structEmbedding struct {
	Container string
	Member    string
	Pos       token.Position
}

// 常量引用
type constRefer struct {
	Name string
	Pos  token.Position
}

// 全局变量引用
type globalRefer struct {
	Name string
	Type string
	Pos  token.Position
}

// 函数调用
type functionCall struct {
	FunName string
	Pos     token.Position
}
