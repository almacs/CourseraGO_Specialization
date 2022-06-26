package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func SplitArray(arr []int, n int) [][]int {
	arrays := [][]int{}

	blockSize := int(math.Ceil(float64(len(arr)) / float64(n)))

	s, e := 0, 0
	for {
		e = Min(s+blockSize, len(arr))
		arrays = append(arrays, arr[s:e])

		if e == len(arr) {
			break
		}
		s = e
	}
	return arrays
}

func SortedMerge(arr [][]int) []int {
	if len(arr) <= 1 {
		return arr[0]
	}

	result := []int{}

	l := arr[0]
	r := arr[1]

	i, j := 0, 0

	for {
		if i > len(l)-1 && j > len(r)-1 {
			break
		}

		if j > len(r)-1 && i <= len(l)-1 {
			result = append(result, l[i])
			i++
			continue
		}

		if i > len(l)-1 && j <= len(r)-1 {
			result = append(result, r[j])
			j++
			continue
		}

		if l[i] < r[j] {
			result = append(result, l[i])
			i++
			continue
		}

		if l[i] > r[j] {
			result = append(result, r[j])
			j++
			continue
		}

		if l[i] == r[j] {
			result = append(result, l[i])
			result = append(result, r[j])
			i++
			j++
			continue
		}
	}

	head := [][]int{result}
	return SortedMerge(append(head, arr[2:]...))
}

func main() {
	//arr := []int{5, 2, 8, 7, 1, 3, 6, 4, 10, 9, 0, -1, 20, 15}

	fmt.Print("enter an array of integer\n")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var arr []int
	tmp := strings.Split(scanner.Text(), " ")

	for _, x := range tmp {
		n, _ := strconv.Atoi(x)
		arr = append(arr, n)
	}

	numTasks := 4
	fmt.Println(arr)

	tasks := SplitArray(arr, numTasks)

	fmt.Println(tasks)

	wg := &sync.WaitGroup{}
	submission := [][]int{}

	fmt.Println("sort concurrently...")
	for _, t := range tasks {
		wg.Add(1)
		go func(a []int) {
			fmt.Printf("sorting subarray %v\n", a)
			sort.Ints(a)
			submission = append(submission, a)
			wg.Done()
		}(t)
	}
	wg.Wait()

	fmt.Println("original tasks:")
	fmt.Println(tasks)
	fmt.Println("task submission:")
	fmt.Println(submission)

	result := SortedMerge(submission)
	fmt.Println("result:")
	fmt.Println(result)
}
