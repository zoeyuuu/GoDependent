package file

import "fmt"

const TIPS = `A`

type Speaker interface {
	Hello1()
	Hello2()
}

type A struct{}

func (a *A) printA() {
	fmt.Println("调用A的print方法")
}
func printNormal() {
	fmt.Println("普通print方法")
}
func printNormal1() {
	fmt.Println("普通print方法")
}
