package main

// 所有文件的绝对路径名
var files []string

func main() {

	getFileName()
	findFileInfos()

	//printFileInfo(info2)
	findDependencies()
	printDenpendencies(dependency)

}
