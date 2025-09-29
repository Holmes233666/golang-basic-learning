package main

import "fmt"

func printArray(nums [10]int) { // 值传递：这一点和C/C++不一样了！
	nums[0] = 100
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	// 固定长度的数组
	var nums [10]int

	nums2 := [10]int{1, 2, 3, 4}

	var nums3 = [4]int{1, 2, 3, 4}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for idx, num := range nums2 {
		fmt.Println(idx, num)
	}

	// 查看数组的数据类型
	fmt.Printf("nums数组的数据类型为：%T\n", nums[0])
	fmt.Printf("nums2数组的数据类型为：%T\n", nums2[0])
	fmt.Printf("nums3数组的数据类型为：%T\n", nums3[0])

	// 尝试用函数printArray打印nums3
	printArray(nums)
	// printArray(nums3) // 报错： cannot use nums3 (variable of type [4]int) as [10]int value in argument to printArray
	fmt.Printf("out of printArry")
	for idx, num := range nums {
		fmt.Println(idx, num)
	}
}
