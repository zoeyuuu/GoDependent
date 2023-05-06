package main

import (
	"GoDependent/file"
	"GoDependent/tool"
	"fmt"
)

func main() {

	file.GetFileName()
	file.FindFileInfos()
	tool.FindDependenyAll()
	info := showFileinfo("context_appengine.go")

	//tool.FindDependencytest("context_appengine.go", "gin.go")

	dep := showDependency("context_appengine.go", "gin.go")
	fmt.Println(info, dep)
}

func showFileinfo(filename string) file.FileInfo {
	for i, v := range file.InfoList {
		if file.InfoList[i].FileBaseName == filename {
			return v
		}
	}
	return file.FileInfo{}
}

func showDependency(file1, file2 string) tool.Dependencies {
	for i, v := range tool.DependencyList {
		if tool.DependencyList[i].Src == file1 && tool.DependencyList[i].Des == file2 {
			return v
		}
	}
	fmt.Println("no dependency")
	return tool.Dependencies{}
}
