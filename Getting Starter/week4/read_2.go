package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	nameSlice := make([]Name, 0)
	fmt.Printf("%s", "Enter the name of the file: ")
	fileName := ""
	_, err := fmt.Scanf("%s", &fileName)
	if err != nil {
		os.Exit(1)
	}

	fd, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fd.Close()
				break
			} else {
				fmt.Println(err)
				fd.Close()
				os.Exit(1)
			}
		}
		line = strings.TrimSuffix(line, "\n")
		lineSlice := strings.Split(line, " ")
		if len(lineSlice) != 2 {
			fmt.Println("Expected a line with two space separated values.")
			fd.Close()
			os.Exit(1)
		}
		firstName := alterStringToCorrectLength(lineSlice[0])
		lastName := alterStringToCorrectLength(lineSlice[1])

		nameSlice = append(nameSlice, Name{firstName, lastName})
	}

	for _, v := range nameSlice {
		fmt.Printf("First name: %s Last name: %s\n", v.fname, v.lname)
	}
}

func alterStringToCorrectLength(s string) string {
	if len(s) > 20 { // If too long string, cuts the string to 20 chars
		runeSlice := []rune(s)
		cutString := string(runeSlice[20:])
		return strings.TrimRight(s, cutString)
	} else { // Pads the string to 20 chars
		return fmt.Sprintf("%20s", s)
	}
}
