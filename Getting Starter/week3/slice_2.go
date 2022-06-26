package main

import (
	"fmt"
	"strconv"
)

func main() {

	theslice := make([]int, 0, 3)
	for true {
		fmt.Print("Enter a integer... ")
		var theinput string
		fmt.Scanln(&theinput)
		if theinput == "X" {
			break
		}
		var intstr *string = &theinput
		nx, err := strconv.Atoi(*intstr)
		if err == nil {
			//if len(theslice) < 3 { <- if i need to truncate the slice size to 3
			theslice = append(theslice, nx)
			fmt.Printf("%d\n", theslice)
			//} <- end of
		}
	}
}
