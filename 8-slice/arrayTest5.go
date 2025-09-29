package main

import "fmt"

func main() {
	nums := []int{1, 2, 3} // len = 3, cap = 3
	s1 := nums[0:2]        // 左闭右开

	// nums和s1其实指向一个数组
	s1[0] = 100
	fmt.Println(nums)

	// copy可以进行深拷贝
	s2 := make([]int, 3)
	copy(s2, nums) // 将nums中的值拷贝到s2中
	fmt.Println(s2)
}
