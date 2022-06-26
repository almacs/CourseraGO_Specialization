package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
	Write a program which prompts the user to first enter a name, and then enter an address.
	Your program should create a map and add the name and address to the map using the keys â€œnameâ€ and â€œaddressâ€,
	respectively. Your program should use Marshal() to create a JSON object from the map, and then your
	program should print the JSON object.
*/

func main() {
	book := make(map[string]string)
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter name: ")
	in.Scan()

	book["name"] = in.Text()

	fmt.Print("Enter address: ")
	in.Scan()

	book["address"] = in.Text()

	bjson, _ := json.Marshal(book)
	fmt.Println(string(bjson))
}
