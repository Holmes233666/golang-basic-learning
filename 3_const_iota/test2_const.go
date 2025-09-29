package main

import "fmt"

// const定义枚举类型
const (
	BEIJING  = 1
	NANJING  = 0
	SHENGHAI = 2
)

// 可以在const()添加一个关键字 iota，每行的iota都会递增1，第一行iota默认值是0

const (
	BANANA = 10 * iota	// 0
	WATERMALEN			// 10
	APPLE				// 20
)

const (
	aa, bb = iota + 1, iota + 2	// 1, 2
	cc, dd
	ee, ff
	gg, hh = iota + 5, iota + 7	// 8, 10
	ll, mm
)

func main() {
	// 常量的定义（只读属性）
	const len int = 3
	fmt.Println(len)
	// len = 10， 常量是不允许修改的

	fmt.Println("BANANA = ", BANANA, "WATERMALEN = ", WATERMALEN, "APPLE = ", APPLE)
	fmt.Println("aa =", aa, "bb =", bb)
	fmt.Println("cc =", cc, "dd =", dd)
	fmt.Println("ee =", ee, "ff =", ff)
	fmt.Println("gg =", gg, "hh =", hh)
	fmt.Println("ll =", ll, "mm =", mm)

	// iota只能配合const一起使用
	// var a int = iota
}
