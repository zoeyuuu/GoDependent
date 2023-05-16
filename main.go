package main

import (
	"GoDependent/file"
	"GoDependent/tool"
	"GoDependent/visualization"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var infoList []file.FileInfo
var depList []tool.Dependencies

func main() {

	file.GetFileName()
	file.FindFileInfos()
	tool.FindDependenyAll()
	depList = tool.DependencyList
	infoList = file.InfoList
	// 可视化
	visualization.JsonVisualization()
	//按照依赖种类转换成json格式
	//tool.DependenciesToJson()
	//按照文件级别转换成json格式
	//depListToJson(depList)
}

func depListToJson(depList []tool.Dependencies) {
	now := time.Now()
	dirName := fmt.Sprintf("outputs/ByFilename/%s", now.Format("2006-01-02_15-04"))
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		panic(err)
	}
	fileName := "depJson.json"
	filePath := filepath.Join(dirName, fileName)
	jsonData, err := json.MarshalIndent(depList, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON data has been written to file %s\n", filePath)
}
