package tool

// 实例化
type InstantiationJSON struct {
	Src        string
	Des        string
	StructName string
	ObjName    string
	Token      string
	Pos        string
}

// 结构体聚合
type StructAggregationJSON struct {
	Src     string
	Des     string
	Whole   string
	Part    string
	ObjName string
	Pos     string
}

// 结构体嵌套
type StructEmbeddingJSON struct {
	Src       string
	Des       string
	Container string
	Member    string
	Pos       string
}

// 常量引用
type ConstReferJSON struct {
	Src  string
	Des  string
	Name string
	Pos  string
}

// 全局变量引用
type GlobalReferJSON struct {
	Src  string
	Des  string
	Name string
	Type string
	Pos  string
}

// 函数调用
type FunctionCallJSON struct {
	Src     string
	Des     string
	FunName string
	Pos     string
}
