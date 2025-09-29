package main

import "fmt"

func printArray2(myArray []int) { // 引用传递
	myArray[0] = 100
	for _, num := range myArray {
		fmt.Println(num)
	}
}

func main() {
	myArray := []int{1, 2, 3, 4, 5} // 长度根据初始化值决定
	printArray2(myArray)
	fmt.Println("----------------------")
	for _, num := range myArray {
		fmt.Println(num)
	}
}
