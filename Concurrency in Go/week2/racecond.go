package main

import (
	"fmt"
	"sync"
)

/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.
Submission: Upload your source code for the program along with your written explanation of race conditions.
*/

/*
	Explanation:
	Race Condition: There are two gorutines. This are sharing "data" and interact with it simultaneously. This are adding the value 1 or 2 respectively to "data" slice.
	I fixed this race condition adding Mutex. It allowed the function for Locking or Unlocking the resources in this case "data" slice..
*/
func main() {
	var data = make([]int, 0, 3)
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	fmt.Println("Race Condition ... ")
	wg.Add(2)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Goroutine 1")
		m.Lock()
		data = append(data, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Goroutine 2")
		m.Lock()
		data = append(data, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(data)
}
