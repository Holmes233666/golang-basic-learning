package lib1

import "fmt"

const a1 int = 10

func Lib1Test(a int) {
	b := a + a1
	fmt.Printf("The sum of a and a1 is: %v\n", b)
}

func init() {
	fmt.Println("lib1.go init")
}

