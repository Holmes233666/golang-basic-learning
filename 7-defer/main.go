package main

import "fmt"

func main() {
	// 按照栈的顺序执行，先压栈的后执行
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("Hello World1")
	fmt.Println("Hello Defer")
}
