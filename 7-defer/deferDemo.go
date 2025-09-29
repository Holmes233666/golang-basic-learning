package main

import "fmt"

func test() int {
	fmt.Println("return test")
	return 0
}

func returnAndDefer() int {
	defer fmt.Println("main end1")
	return test()
}

func main() {
	returnAndDefer()
}
