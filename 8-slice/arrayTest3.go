package main

import "fmt"

// 切片声明的4种方式
func main() {
	// 1. 声明但不初始化，访问/赋值元素会越界
	var nums1 []int
	fmt.Printf("len = %d, nums1: %v\n", len(nums1), nums1)

	// 2. 声明且初始化
	var nums2 = []int{1, 2, 3, 4}
	fmt.Printf("len = %d, nums2: %v\n", len(nums2), nums2)

	// 3. 使用make进行初始化
	var nums3 = make([]int, 4)
	fmt.Printf("len = %d, nums3: %v\n", len(nums3), nums3)

	// 4. 使用 := 进行推导
	nums4 := make([]int, 5)
	fmt.Printf("len = %d, nums4: %v\n", len(nums4), nums4)

	// 判断一个slice是否为空
	if nums1 == nil {
		fmt.Println("nums1 是一个空切片")
	} else {
		fmt.Println("nums1 不是一个空切片")
	}

}
