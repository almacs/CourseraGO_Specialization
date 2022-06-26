package main

import "fmt"

func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func decFn(x int) int { return x - 1 }

func main() {
	//Variadic arguments
	fmt.Println(getMax(1, 2, 3, 4))
	vslice := []int{1, 2, 3, 4}
	fmt.Println(getMax(vslice...))

	//Deferred Call Arguments
	i := 1
	defer fmt.Println(i + 1)
	i++
	fmt.Println("Hello")
	fmt.Println("task2")

	//Variables as FUnctions
	funcVar = incFn
	fmt.Println("Variable as Funciton: ", funcVar(1))

	//Function as Arguments
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))

	//Anonymous Funcitons
	v := applyIt(func(x int) int { return x + 1 }, 2)
	fmt.Println("Anonymous Functions", v)

}
