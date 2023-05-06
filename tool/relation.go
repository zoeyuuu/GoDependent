package tool

// 实例化
type Instantiation struct {
	TypeName string
	VarName  string
}

// 结构体聚合
type structAggregation struct {
	whole   string
	part    string
	varName string
}

// 结构体嵌套
type structEmbedding struct {
	container string
	member    string
}

// 常量引用
type constRefer struct {
	name string
}

// 全局变量引用
type globalRefer struct {
	name string
	Type string
}
