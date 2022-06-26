package main

import "fmt"

func main() {
	/*
				x := []int{4, 8, 5}
				y := -1

				for _, elt := range x {
					if elt > y {
						y = elt
					}
				}
				fmt.Print(y)

			x := [...]int{4, 8, 5}
			y := x[0:2]
			z := x[1:3]
			y[0] = 1
			z[1] = 3
			fmt.Print(x)
		x := [...]int{1, 2, 3, 4, 5}
		y := x[0:2]
		z := x[1:4]
		fmt.Print(len(y), cap(y), len(z), cap(z))
	*/

	x := map[string]int{
		"ian": 1, "harris": 2}
	for i, j := range x {
		if i == "harris" {
			fmt.Print(i, j)

		}
	}
}
