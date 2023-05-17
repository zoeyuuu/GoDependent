package visualization

import (
	"GoDependent/file"
	"GoDependent/tool"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

var infoList []file.FileInfo
var depList []tool.Dependencies

type node struct {
	Id        string `json:"id"`
	Name      string `json:"relName"`
	BaseName  string `json:"name"`
	Weight    int    `json:"symbolSize"`
	PkgName   string `json:"category"`
	Draggable bool   `json:"draggable"`
}
type edge struct {
	Source     string         `json:"source"`
	Target     string         `json:"target"`
	Relations  map[string]int `json:"relations"`
	Weight     int            `json:"weight"`
	Value      float64        `json:"value"`
	SourceName string         `json:"sourcename"`
	TargetName string         `json:"targetname"`
}
type Category struct {
	Name string `json:"name"`
}

func JsonVisualization() {
	infoList = file.InfoList
	depList = tool.DependencyList
	edges := edgesJson(depList)
	nodes := nodesJson(infoList, edges)
	categories := getCategories(nodes)

	// 创建一个映射，用于根据节点名称查找对应的节点ID
	nodeIDMap := make(map[string]string)
	for _, n := range nodes {
		nodeIDMap[n.Name] = n.Id
	}
	// 将 "source" 和 "target" 的值改为对应的节点ID，符合echarts格式
	for i := range edges {
		edg := &edges[i]
		if id, ok := nodeIDMap[edg.Source]; ok {
			edg.Source = id
		}
		if id, ok := nodeIDMap[edg.Target]; ok {
			edg.Target = id
		}
	}

	data := struct {
		Nodes      []node     `json:"nodes"`
		Edges      []edge     `json:"links"`
		Categories []Category `json:"categories"`
	}{
		Nodes:      nodes,
		Edges:      edges,
		Categories: categories,
	}

	dirName := "visualization/static"
	fileName := "visual.json"
	filePath := filepath.Join(dirName, fileName)
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("FunctionCall_JSON data has been written to file %s\n", filePath)
}

func getCategories(nodes []node) []Category {
	categories := make(map[string]struct{})
	for _, node := range nodes {
		if _, ok := categories[node.PkgName]; !ok {
			categories[node.PkgName] = struct{}{}
		}
	}

	categoryList := []Category{}
	for categoryName := range categories {
		categoryList = append(categoryList, Category{Name: categoryName})
	}
	return categoryList
}

func nodesJson(infoList []file.FileInfo, edges []edge) []node {
	var nodes []node
	for i, info := range infoList {
		weight := 0
		for _, edg := range edges {
			if edg.Source == info.FileRelName || edg.Target == info.FileRelName {
				weight++
			}
		}
		tmp := node{
			Id:        strconv.Itoa(i),
			Name:      info.FileRelName, //相对路径名
			BaseName:  info.FileBaseName,
			Weight:    weight,
			PkgName:   info.PkgName,
			Draggable: true,
		}
		nodes = append(nodes, tmp)
	}
	return nodes
}

func edgesJson(depList []tool.Dependencies) []edge {
	var edges []edge

	// 权重分配
	mp := map[string]int{
		"ConstRefer":        1,
		"GlobalRefer":       1,
		"FunctionCall":      1,
		"instantiation":     1,
		"structEmbedding":   1,
		"structAggregation": 1,
	}

	for _, dep := range depList {
		relations := make(map[string]int)
		for key, values := range dep.Relations {
			relations[key] = len(values)
		}
		var weight int
		for key, value := range relations {
			weight += value * mp[key]
		}
		value := smoothWeight(float64(weight))
		tmp := edge{
			Source:     dep.RelativeSrc,
			Target:     dep.RelativeTar,
			Relations:  relations,
			Weight:     weight,
			Value:      value,
			SourceName: dep.RelativeSrc,
			TargetName: dep.RelativeTar,
		}
		edges = append(edges, tmp)
	}
	return edges
}
