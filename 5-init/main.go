package main

import (
	mylib "github.com/Holmes233666/golang-basic-learning/5-init/lib1" // 一定要写在相对于gopath的相对路径下
	_ "github.com/Holmes233666/golang-basic-learning/5-init/lib2"     // 匿名导包，无法使用该包的方法，但是会执行当前包的init函数
	// . "github.com/Holmes233666/golang-basic-learning/5-init/lib2" // 直接使用，不用使用包.的方式。将fmt包中的全部方法，导入到当前本包的作用中
)

func main() {
	a := 10
	mylib.Lib1Test(a)
	//lib2.Lib2Test(a)
}
