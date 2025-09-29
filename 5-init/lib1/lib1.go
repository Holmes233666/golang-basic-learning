package lib1

import "fmt"

func Lib1Test() {
	fmt.Println("当前lib1提供的接口")
}

func init() {
	fmt.Println("lib1 init")
}
