package main

import "fmt"

func foo(y int) {
	y = y + 1
}

func fooP(y *int) {
	*y = *y + 1
}

func main() {
	x := 2
	foo(x)
	fmt.Println(x)

	fooP(&x)
	fmt.Println(x)
}
