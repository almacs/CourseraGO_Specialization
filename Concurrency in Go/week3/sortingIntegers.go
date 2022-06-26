package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
	Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

	The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

	Submission: Upload your source code for the program.
*/

func sortQuarter(c []int, sS chan []int) {
	fmt.Println("subarray to sort:", c)
	sort.Ints(c)
	sS <- c
}

func main() {
	var inputStr string
	//Ask for the serie
	fmt.Println("Please enter series of integers separeted by comma:")
	fmt.Scanln(&inputStr)

	serie := make([]int, 0, 3)

	var ns []string = strings.Split(inputStr, ",")

	for _, s := range ns {
		n, e := strconv.Atoi(s)
		if e != nil {
			fmt.Println("The serie is not allowed. Try again.")
			break
		}

		serie = append(serie, n)
	}
	//create a channel
	var serieSort = make(chan []int)
	l := len(ns) / 4
	j := 0
	r := len(ns) % 4

	for x := 0; x < 4; x++ {
		//Calling goroutine with 1/4 of array, if the size is not divided in 4, I'm adding the remainder in the last subarray
		go sortQuarter(serie[j:(j+l)+int(x/3*r)], serieSort)
		j = j + l
	}
	//Receiving data from channel
	seriea := <-serieSort
	serieb := <-serieSort
	seriec := <-serieSort
	seried := <-serieSort
	seriea = append(seriea, serieb...)
	seriea = append(seriea, seriec...)
	seriea = append(seriea, seried...)
	//if we need all the serie sorted we have to uncomment the next line
	//sort.Ints(seriea)

	//Printing the entire sorted list
	fmt.Println("Entire sorted list:", seriea)
}
