package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Enter list of integers (space delimiter): ")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	if reader.Err() != nil {
		fmt.Printf("%s \n", reader.Err())
		return
	}

	ints, err := ToInt(strings.Split(reader.Text(), " "))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n ", *BubbleSort(&ints))
}

func ToInt(strs []string) ([]int, error) {
	ints := make([]int, 0, len(strs))
	for _, str := range strs {
		j, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		ints = append(ints, j)
	}
	return ints, nil
}

func BubbleSort(intps *[]int) *[]int {
	ints := *intps
	var upto = len(ints) - 1
	for swapped := true; swapped; {
		swapped = false
		for i := 0; i < upto; i = i + 1 {
			if ints[i] > ints[i+1] {
				Swap(&ints, i)
				swapped = true
			}
		}
		upto = upto - 1
	}
	return intps
}

func Swap(intps *[]int, index int) {
	ints := *intps
	temp := ints[index]
	ints[index] = ints[index+1]
	ints[index+1] = temp
}
