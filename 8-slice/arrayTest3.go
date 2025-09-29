package main

import (
	"fmt"
)

func main() {
	// slice的声明方式：
	// 方式1：声明 + 定义
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("当前slice1的长度为", len(slice1))

	// 方式2：只声明
	slice2 := []int{}
	slice2 = make([]int, 3)    // 开辟空间
	slice2 = append(slice2, 1) // 追加一个元素
	fmt.Println("当前slice2的长度为", len(slice2))

	// 方式3：声明 + 初始化 [方式1 + 方式2 的合并方式]
	var slice3 = make([]int, 4)
	fmt.Println(slice3)

	// 方式4：直接使用推断
	slice4 := make([]int, 5)
	fmt.Println(slice4)

	var slice5 []int
	if slice5 == nil {
		fmt.Println("slice5 is nil")
	} else {
		fmt.Println("slice5 is not nil")
	}

}
