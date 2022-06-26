package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var fileName string

	fmt.Println("Enter text file name:")
	fmt.Scanln(&fileName)

	dat, e := os.Open(fileName)
	check(e)

	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanLines)

	sl := make([]name, 0, 2)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")

		if len(str) != 2 {
			continue
		}
		p1 := name{str[0], str[1]}

		sl = append(sl, p1)

	}

	dat.Close()

	for _, n := range sl {
		fmt.Println(n.fname, n.lname)
	}

}
