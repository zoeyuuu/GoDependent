package main

import "fmt"

func main() {

	getFileName()
	findFileInfos()
	//findDependency(1, 0)
	findDependenyAll()

	fmt.Println(dependencyList)
}
