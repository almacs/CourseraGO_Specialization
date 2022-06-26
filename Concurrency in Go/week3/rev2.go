package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Enter integer Numbers:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := strings.Fields(scanner.Text())

	arr := make([]int, 0)
	for _, s := range str {
		v, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("can not convert %s to int: %v\n", s, err)
		}
		arr = append(arr, v)
	}

	wg := sync.WaitGroup{}
	wg.Add(4)
	a := 0
	inc := len(arr) / 4
	b := inc
	for i := 0; i < 3; i++ {
		go sortArray(arr[a:b], &wg)
		a = b
		b += inc
	}
	go sortArray(arr[a:], &wg)
	wg.Wait()
	sort.Ints(arr)
	fmt.Println("Final sorted array:", arr)
}

func sortArray(arr []int, wg *sync.WaitGroup) {
	fmt.Println("Going to Sort:", arr)
	sort.Ints(arr)
	wg.Done()
}

/*

	Output:

	Enter integer Numbers:
	3 5 7 8 9 0 18 299 4 7 980 3
	Going to Sort: [7 980 3]
	Going to Sort: [8 9 0]
	Going to Sort: [3 5 7]
	Going to Sort: [18 299 4]
	Final sorted array: [0 3 3 4 5 7 7 8 9 18 299 980]

*/
