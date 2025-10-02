package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[0:2] // 左闭右开
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

	s1[0] = 5
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 4)
	// s1内容拷贝到s3中
	copy(s3, s1)
}
