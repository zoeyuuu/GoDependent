package file

import "fmt"

const TIP = `A`
const TIP2 = `A`

type B struct {
	a3 A
}

var a2 A

func a() {
	var a2 A
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
