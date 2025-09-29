package main

import (
	_ "GolangStudy/5-init/lib1"
	myLib2 "GolangStudy/5-init/lib2"
	// . "GolangStudy/5-init/lib2"	// 直接导入到当前包，不用加包名再引用函数
)

func main() {
	myLib2.Lib2Test()
}
