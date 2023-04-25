package main

var infoList []fileInfo
var files []string

func main() {

	getFileName()

	// filename1 := "./file/b.go"
	filename1 := files[0]
	filename2 := "./file/a.go"
	println(filename1, filename2)
	info1 := findFileInfo(filename1)
	info2 := findFileInfo(filename2)

	infoList = append(infoList, info1, info2)
	//printFileInfo(info2)
	findDependencies()
	printDenpendencies(dependency)

}
