package main

import (
	"fmt"
)

// 值拷贝形式的传参
func printArray(a [4]int) { // 注意类型匹配问题
	for i := 0; i < len(a); i++ {
		fmt.Println("index:", i, ", val = ", a[i])
	}
	fmt.Println("-------------------------------")
}

func main() {
	// 声明一个长度为10的数组，并进行遍历
	var nums1 [10]int
	for i := 0; i < len(nums1); i++ {
		fmt.Println("index: ", i, "num: ", nums1[i])
	}
	fmt.Println("-------------------------------")
	// 声明一个长度为10，初始化部分数据的数组，并进行遍历
	var nums2 = [10]int{1, 2, 3, 4} // 初始化了，需要用=
	for idx, val := range nums2 {
		fmt.Println("index: ", idx, "num: ", val)
	}
	fmt.Println("-------------------------------")
	nums3 := [4]int{1, 2, 3, 4}
	for idx, val := range nums3 {
		fmt.Println("index: ", idx, "num: ", val)
	}
	fmt.Println("-------------------------------")
	// 查看三个数组的类型
	fmt.Printf("The type of nums1 is: %T\n", nums1)
	fmt.Printf("The type of nums2 is: %T\n", nums2)
	fmt.Printf("The type of nums3 is: %T\n", nums3)
	fmt.Println("-------------------------------")
	printArray(nums3)

}
