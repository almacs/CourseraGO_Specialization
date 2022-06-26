package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return ((a * math.Pow(t, 2)) / 2) + (v * t) + s
	}
	return fn
}

func main() {
	var a float64
	var vo float64
	var so float64
	var t float64

	fmt.Println("Enter a accelerationvalue:")
	fmt.Scanln(&a)
	fmt.Println("Enter initial velocity:")
	fmt.Scanln(&vo)
	fmt.Println("Enter initial displacement:")
	fmt.Scanln(&so)

	fmt.Println("Enter a value for time")
	fmt.Scanln(&t)

	fn := GenDisplaceFn(a, vo, so)
	fmt.Println(fn(t))

}
