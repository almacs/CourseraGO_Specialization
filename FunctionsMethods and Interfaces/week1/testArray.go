package main

import "fmt"

func foo(x [3]int) int {
	return x[0]
}

func fooP(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}

func fooS(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(foo(a))

	fooP(&a)
	fmt.Println(a)

	a2 := []int{1, 2, 3}
	fooS(a2)
	fmt.Println(a2)
}
