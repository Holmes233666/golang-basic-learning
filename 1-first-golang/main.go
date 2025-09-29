package main	// 包名

import (
	"fmt"	// 导入 fmt 包
	"time"	
)

func main() {	// 左花括号 “{” 一定和函数名同行，不可另起一行
	fmt.Println("Hello, World!")	// 没有分号，语句结束由换行符决定；保持不加分号的习惯
	time.Sleep(2 * time.Second)
}