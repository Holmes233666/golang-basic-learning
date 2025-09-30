package lib2

import "fmt"

const b int = 20

func Lib2Test(a2 int) {
	c := a2 + b
	fmt.Printf("The sum of a2 and b is %v\n", c)
}

func init() {
	fmt.Println("lib2.go init")
}