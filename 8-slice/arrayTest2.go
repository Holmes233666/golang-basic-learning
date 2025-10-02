package main

import "fmt"

// 引用传递
func printArray2(a []int) {
	//for i := 0; i < len(a); i++ {
	//	fmt.Printf("%d ", a[i])
	//}
	for _, val := range a {
		fmt.Print(val, " ")
	}
	fmt.Print("\n") // 不加这一行，zsh会自动补上一个%
	a[0] = 199
}

func main() {
	// 动态数组
	nums1 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums1); i++ {
		fmt.Printf("index %d: %d\n", i, nums1[i])
	}
	// 类型打印
	fmt.Printf("The type of nums1 is: %T\n", nums1) // []int 就是其类型

	// 使用函数进行打印
	printArray2(nums1)
	// 第二次打印
	printArray2(nums1)
}
