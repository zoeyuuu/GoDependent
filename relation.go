package main

// 实例化
type instantiation struct {
	typeName string
	varName  string
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
