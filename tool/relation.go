package tool

import "go/token"

// 实例化
type Instantiation struct {
	TypeName string
	VarName  string
	pos      token.Position
}

// 结构体聚合
type structAggregation struct {
	whole   string
	part    string
	varName string
	pos     token.Position
}

// 结构体嵌套
type structEmbedding struct {
	container string
	member    string
	pos       token.Position
}

// 常量引用
type constRefer struct {
	name string
	pos  token.Position
}

// 全局变量引用
type globalRefer struct {
	name string
	Type string
	pos  token.Position
}

// 函数调用
type functionCall struct {
	funName string
	pos     token.Position
}
