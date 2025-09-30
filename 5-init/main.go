package main

import (
	lib1 "github.com/Holmes233666/golang-basic-learning/5-init/lib1"	// 一定要写在相对于gopath的相对路径下
	lib2 "github.com/Holmes233666/golang-basic-learning/5-init/lib2"
)

func main() {
	a := 10
	lib1.Lib1Test(a)
	lib2.Lib2Test(a)
}