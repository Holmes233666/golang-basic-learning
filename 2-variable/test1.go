package main


import "fmt"

// 声明全局变量时，必须使用 var 关键字，【不能使用 := 方式】
var mm int

// go中变量的四种声明方式
func main() {
	fmt.Printf("type of mm is %T, value of mm is %v\n", mm, mm)
	// 1. var 变量名 变量类型
	// 声明一个整型变量a，默认值为0
	var a int
	fmt.Printf("type of a is %T, value of a is %v\n", a, a)

	// 2. var 变量名 变量类型 = 初始值
	// 声明一个整型变量b，初始值为10
	var b int = 10
	fmt.Printf("type of b is %T, value of b is %v\n", b, b)

	// 3. var 变量名 = 初始值	【省去数据类型】
	// 声明一个整型变量c，初始值为20
	var c = 10
	fmt.Printf("type of c is %T, value of c is %v\n", c, c)


	// 声明一个字符串变量d，初始值为"hello"
	var s = "hello"
	fmt.Printf("type of s is %T, value of s is %v\n", s, s)

	// 4. 变量名 := 初始值	【只能在函数体内使用，省去var和数据类型】
	// 声明一个整型变量e，初始值为30
	e := 30
	fmt.Printf("type of e is %T, value of e is %v\n", e, e)

	// 声明多个变量
	var x, y int = 100, 200
	fmt.Printf("type of x is %T, value of x is %v\n", x, x)
	fmt.Printf("type of y is %T, value of y is %v\n", y, y)

	var m, n = 300, "hello"
	fmt.Printf("type of m is %T, value of m is %v\n", m, m)
	fmt.Printf("type of n is %T, value of n is %v\n", n, n)

	p, q := 400, "world"
	fmt.Printf("type of p is %T, value of p is %v\n", p, p)
	fmt.Printf("type of q is %T, value of q is %v\n", q, q)

	// 多行写法
	var (
		NODE_NAME = "node1"
		NODE_IP   = "192.168.1.1"
	)
	fmt.Println("value of NODE_NAME is", NODE_NAME, ", vlaue of NODE_IP is", NODE_IP)
}