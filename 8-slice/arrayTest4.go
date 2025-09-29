package main

import "fmt"

func main() {
	// 设置len = 3，cap = 5的slice
	slice := make([]int, 3, 5)
	fmt.Printf("slice:%v, cap = %d, len = %d\n", slice, cap(slice), len(slice)) // slice:[0 0 0], cap = 5, len = 3

	// 向slice中追加一个元素
	slice = append(slice, 1)
	fmt.Printf("slice:%v, cap = %d, len = %d\n", slice, cap(slice), len(slice))

	// 再追加一个元素，此时已经满了
	slice = append(slice, 2)
	fmt.Printf("slice:%v, cap = %d, len = %d\n", slice, cap(slice), len(slice))

	fmt.Printf("---------------------\n")

	slice = append(slice, 3)
	fmt.Printf("slice:%v, cap = %d, len = %d\n", slice, cap(slice), len(slice))

	// 如果没有设置cap，那么默认cap = len
	nums := make([]int, 3)
	fmt.Printf("nums:%v, cap = %d, len = %d\n", nums, cap(nums), len(nums))
	nums = append(nums, 1)
	fmt.Printf("nums:%v, cap = %d, len = %d\n", nums, cap(nums), len(nums))
}
