package main

import (
	"GoDependent/file"
	"GoDependent/tool"
	"GoDependent/visualization"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var infoList []file.FileInfo
var depList []tool.Dependencies

func main() {

	root := "D:\\projects\\go_projects\\src\\ginVersions\\gin-1.9.0"
	file.GetFileName(root)
	file.FindFileInfos()
	tool.FindDependenyAll()
	depList = tool.DependencyList
	infoList = file.InfoList

	//按照依赖种类转换成json格式
	tool.DependenciesToJson()

	//按照文件级别转换成json格式
	depListToJson(depList)

	// 数据处理
	visualization.JsonVisualization()
	//relationGraphShow()
}

func depListToJson(depList []tool.Dependencies) {
	now := time.Now()
	dirName := fmt.Sprintf("outputs/"+file.Name+"/ByFilename/%s", now.Format("2006-01-02_15-04"))
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

func relationGraphShow() {
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/static", "visualization")

	// 设置HTML文件路径
	filePath := "visualization/test.html"

	r.GET("/", func(c *gin.Context) {
		c.File(filePath)
	})

	// 启动HTTP服务器
	fmt.Println("Server is running on http://localhost:8081")
	log.Fatal(r.Run(":8081"))
}
