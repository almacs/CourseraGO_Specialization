package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(nums []int) {
	sorted := false
	for sorted == false {
		sorted = true
		for i := 0; i+1 < len(nums); i++ {
			if nums[i] > nums[i+1] {
				Swap(nums, i)
				sorted = false
			}
		}
	}
}

func Swap(nums []int, idx int) {
	if len(nums) > idx+1 {
		temp := nums[idx]
		nums[idx] = nums[idx+1]
		nums[idx+1] = temp
	}
}

func GetNumsFromUser() []int {
	nums := make([]int, 0, 10)
	good_input := false
	for good_input == false {
		fmt.Printf("Enter a list of up to 10 ints to sort, separated by commas (','): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input_string := scanner.Text()
		input_string = strings.ReplaceAll(input_string, " ", "")
		num_strings := strings.Split(input_string, ",")
		nums = make([]int, 0, len(num_strings))

		for _, num_string := range num_strings {
			num, _ := strconv.Atoi(num_string)
			nums = append(nums, num)
		}

		if len(nums) > 10 {
			fmt.Println("\ntoo many integers.  please enter a list of up to 10 integers...\n")
		} else {
			good_input = true
		}
	}

	return nums
}

func PrintNums(nums []int) {
	fmt.Printf("Sorted list:  ")
	for idx, num := range nums {
		fmt.Printf("%d", num)
		if idx+1 == len(nums) {
			fmt.Println()
		} else {
			fmt.Printf(",")
		}
	}
}

func main() {
	numbers := GetNumsFromUser()
	BubbleSort(numbers)
	PrintNums(numbers)
}
