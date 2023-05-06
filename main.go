package main

import (
	"GoDependent/file"
	"GoDependent/tool"
	"fmt"
)

var infoList []file.FileInfo
var depList []tool.Dependencies
var fileDepList []tool.Dependencies

func main() {

	file.GetFileName()
	file.FindFileInfos()
	tool.FindDependenyAll()
	//info := showFileinfo("context_appengine.go")

	depList = tool.DependencyList
	infoList = file.InfoList

	findDepOfFile("gin.go")

	dep := showDependency("context_appengine.go", "gin.go")
	fmt.Println(depList, dep)
}

func findDepOfFile(filename string) {
	for _, dep := range depList {
		if dep.Src == filename || dep.Des == filename {
			fileDepList = append(fileDepList, dep)
		}
	}
}

func showFileinfo(filename string) file.FileInfo {
	for i, v := range file.InfoList {
		if infoList[i].FileBaseName == filename {
			return v
		}
	}
	return file.FileInfo{}
}

func showDependency(file1, file2 string) tool.Dependencies {
	for i, v := range depList {
		if depList[i].Src == file1 && depList[i].Des == file2 {
			return v
		}
	}
	fmt.Println("no dependency")
	return tool.Dependencies{}
}
