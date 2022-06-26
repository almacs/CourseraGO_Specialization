package main

import (
	"fmt"
	"strconv"
)

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, i int) {
	temp := sli[i+1]
	sli[i+1] = sli[i]
	sli[i] = temp
}

func main() {
	a := make([]int, 0, 10)
	var n string

	fmt.Println("Enter a sequence of up to 10 integers...")
	i := 0
	for i < 10 {
		fmt.Scanln(&n)
		num, e := strconv.Atoi(n)
		if e != nil {
			fmt.Println("Enter a valid integer")
			continue
		}
		a = append(a, num)
		i = i + 1
	}

	fmt.Println("Original sequence:", a)
	BubbleSort(a)
	fmt.Println("Sequence sorted: ", a)
}
