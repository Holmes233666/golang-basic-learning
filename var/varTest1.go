package main

import "fmt"

// 声明全局变量，方法1-3是可以的

var g_A int = 100
var g_B = 200

// := 只能用在函数体内声明 + 定义一个变量
//g_C := 300， 会报错
/*
四种变量的声明方式
*/
func main() {
	// -------------------单变量的声明----------------------
	// 方法1：声明一个变量，默认的值是 0
	var a int
	fmt.Printf("type of a: %T\n", a)

	// 方法2：声明一个变量，给出初始化值
	var b int = 100
	fmt.Printf("type of b: %T\n", b)

	// 方法3：在初始化的时候省去对变量类型的说明，根据值自动匹配当前的数据类型
	var c = 1000
	fmt.Printf("type of c: %T\n", c)

	var test_str string = "hello world"
	fmt.Printf("type of %s: %T\n", test_str, test_str)

	var test_str2 = "hello world"
	fmt.Printf("type of %s: %T\n", test_str2, test_str2)

	// 方法4：【最常用的写法】省去var关键字，直接自动匹配
	e := 100 // 既声明， 又初始化
	fmt.Printf("type of %s: %T\n", e, e)

	g := 3.14
	fmt.Printf("type of %s: %T\n", g, g)

	s := "string s"
	fmt.Printf("type of %s: %T\n", s, s)

	fmt.Println("g_A = ", g_A, "g_B = ", g_B)

	// -------------------多变量的声明----------------------
	var aa, bb int = 1, 2
	fmt.Println("aa = ", aa, "bb = ", bb)
	var num, str = 100, "hello world"
	fmt.Println("num = ", num, "str = ", str)

	// 多行的多变量声明
	var (
		am = 1
		bm = "boolean type"
	)
	fmt.Println("am = ", am, "bm = ", bm)
}
