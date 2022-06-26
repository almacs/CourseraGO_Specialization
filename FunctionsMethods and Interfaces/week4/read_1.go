package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Names struct {
	fname string
	lname string
}

func main() {
	var file string
	fmt.Println("Enter filename:")
	_, err := fmt.Scan(&file)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	namesColl := make([]Names, 0, 50)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		np := new(Names)
		ln := scanner.Text()
		names := strings.Fields(ln)
		np.fname = names[0]
		np.lname = names[1]
		namesColl = append(namesColl, *np)
	}

	fmt.Println("The file contains the following names:")
	for _, v := range namesColl {
		fmt.Printf("Firstname: %s, Lastname: %s\n", v.fname, v.lname)
	}
}
