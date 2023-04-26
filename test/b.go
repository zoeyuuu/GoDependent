package main

import (
	"fmt"
	"test/file"
	"test/file/file1"
)

const TIP = `A`
const TIP2 = `A`

type B struct {
	a3 A
}

var a2, a3 A
var c1, c2 file.C

func a() {
	file.PrintNormalC()
	file1.PrintNormalD()
	a1 := A{}
	a2.printA()
	a1.printA()
	printNormal()
	printNormal()
	printNormal1()
	fmt.Println(TIPS)
}

func (b *B) printB() {
	fmt.Println("调用A的print方法")
}
func printNormalB() {
	fmt.Println("普通print方法")
}
