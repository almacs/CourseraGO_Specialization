package main

import (
	"fmt"
)

var i int = 0

//var wg sync.WaitGroup

func inc2() {
	i = i + 1
	//wg.Done()
}
func inc() {
	i = 4
	//wg.Done()
}
func main() {
	//wg.Add(2)
	go inc()
	go inc2()
	//wg.Wait()
	fmt.Println(i)
}
