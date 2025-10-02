package main

import "fmt"

// 切片的追加
func main() {
	// capacity 与 len 的区别
	nums1 := make([]int, 3, 5) // len = 3, cap = 5
	fmt.Printf("len(nums1) = %d, capacity(nums1) = %d, nums1 = %v\n", len(nums1), cap(nums1), nums1)

	// 元素追加
	nums1 = append(nums1, 1)
	fmt.Printf("len(nums1) = %d, capacity(nums1) = %d, nums1 = %v\n", len(nums1), cap(nums1), nums1) // len(nums1) = 4, capacity(nums1) = 5, nums1 = [0 0 0 1]
	nums1 = append(nums1, 2)
	fmt.Printf("len(nums1) = %d, capacity(nums1) = %d, nums1 = %v\n", len(nums1), cap(nums1), nums1) // len(nums1) = 5, capacity(nums1) = 5, nums1 = [0 0 0 1 2]
	// cap已满，会再追加为原cap的2倍： cap = cap * 2
	nums1 = append(nums1, 3)
	fmt.Printf("len(nums1) = %d, capacity(nums1) = %d, nums1 = %v\n", len(nums1), cap(nums1), nums1) // len(nums1) = 6, capacity(nums1) = 10, nums1 = [0 0 0 1 2 3]

	// 如果不设置cap， cap = len
	nums2 := make([]int, 3)
	fmt.Printf("len(nums2) = %d, capacity(nums2) = %d, nums2 = %v\n", len(nums2), cap(nums2), nums2) // len(nums2) = 3, capacity(nums2) = 3, nums2 = [0 0 0]
}
