package main

import "fmt"

func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())
	fmt.Print(fB())

	i := 1

	fmt.Print(i)

	i++

	defer fmt.Print(i + 1)

	fmt.Print(i)
}
