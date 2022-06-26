package main

import "fmt"

func foo(x int, y int) {
	fmt.Println(x * y)
}

func foo2(x int) int {
	return x + 1
}

func foo3(x int) (int, int) {
	return x, x + 1
}

func main() {
	foo(2, 3)

	y := foo2(2)
	fmt.Println(y)

	a, b := foo3(3)
	fmt.Println(a, b)

}
