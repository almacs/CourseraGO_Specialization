package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var sInput []int

	sInput = make([]int, 0, 3)
	//sLen := len(sInput)
	//x := 0

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an integer or 'X' for exiting:")

	for scanner.Scan() {
		st := scanner.Text()
		if st == "X" {
			break
		}

		intVar, err := strconv.Atoi(st)
		if err != nil {
			fmt.Println("Enter a valid integer or 'X' for exiting")
			continue
		}

		//if x < sLen {
		//	sInput[x] = int(intVar)
		//} else {
		sInput = append(sInput, int(intVar))
		sort.Ints(sInput)
		//}
		//x = x + 1
		fmt.Println("slice:", sInput)
	}

	sort.Ints(sInput)
	fmt.Println("slice:", sInput)
}
