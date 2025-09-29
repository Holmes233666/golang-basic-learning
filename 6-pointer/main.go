package main

import "fmt"

func myswap(a *int, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

func main() {
	a := 10
	b := 20
	myswap(&a, &b)
	fmt.Println("a=", a, "b=", b)
}
