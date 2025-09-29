package main

import "fmt"

func main() {
	// 声明myMap为key为int，value为string
	var myMap map[int]string // key是int, value是string
	if myMap == nil {
		fmt.Println("map is nil")
	}
	// 使用map前必须得初始化空间大小
	myMap = make(map[int]string, 10)
	myMap[1] = "hello"
	myMap[2] = "world"

	fmt.Println("myMap len is ", len(myMap), ", the content is", myMap)

	// 第二种声明方式
	myMap2 := make(map[string]string, 10)
	myMap2["hello"] = "world"
	fmt.Println("myMap2 len is ", len(myMap2), ", the content is", myMap2)

	// 第3种初始化方式
	myMap3 := map[string]string{
		"one": "cpp",
	}
	fmt.Println("myMap3 len is ", len(myMap3), ", the content is", myMap3)

}
