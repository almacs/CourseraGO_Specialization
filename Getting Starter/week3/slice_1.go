package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var str string
	var sli []int
	sli = make([]int, 0, 3)

	for {

		fmt.Print("\n\nFor exit the program, write X \n")

		fmt.Scan(&str)
		if "X" != str {
			i, err := strconv.Atoi(str)
			if err == nil {
				sli = append(sli, i)
				sort.Ints(sli)
			}
			fmt.Print("\n")
			fmt.Print(sli)
		} else {
			break
		}
	}
}
