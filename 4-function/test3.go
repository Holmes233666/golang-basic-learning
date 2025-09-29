package main

import "fmt"

// 单返回值
func tool(a string, b int) int {
	fmt.Println("--------tool1--------")
	fmt.Println("a:", a, "b:", b)
	c := 100
	return c
}

// 多返回值，匿名
func tool2(a string, b int) (int, int) {
	fmt.Println("--------tool2--------")
	fmt.Println("a:", a, "b:", b)
	c := b + 100
	d := b + 1000
	return c, d
}

// 多返回值，起名
func tool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("--------tool3--------")
	fmt.Println("a:", a, "b:", b)
	fmt.Println("r1, r2是属于tool3的形参，默认初始化值是0")
	fmt.Println("r1:", r1, "r2:", r2)
	r1 = b + 1000
	r2 = b + 100
	return
}

// 多返回值，相同类型合并
func tool4(a string, b int) (r1, r2 int) {
	fmt.Println("--------tool4--------")
	fmt.Println("a:", a, "b:", b)
	r1 = b + 1000
	r2 = b + 100
	return
}

func main() {
	c := tool("abc", 555)
	fmt.Println(c)

	ret1, ret2 := tool2("abc", 100)
	fmt.Println(ret1, ret2)

	r1, r2 := tool3("abc", 100)
	fmt.Println(r1, r2)

	r1, r2 = tool4("abc", 100)
	fmt.Println(r1, r2)
}
