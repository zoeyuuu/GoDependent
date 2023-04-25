package main

// 111
var infoList []fileInfo

func main() {

	filename1 := "./file/b.go"
	filename2 := "./file/a.go"
	info1 := findFileInfo(filename1)
	info2 := findFileInfo(filename2)

	infoList = append(infoList, info1, info2)
	//printFileInfo(info2)

	findDependencies()
	printDenpendencies(dependency)
}
