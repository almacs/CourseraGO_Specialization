package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (c Cow) Eat() {
	println("grass")
}

func (c Cow) Move() {
	println("walk")
}

func (c Cow) Speak() {
	println("moo")
}

type Bird struct {
}

func (b Bird) Eat() {
	println("worms")
}

func (b Bird) Move() {
	println("fly")
}

func (b Bird) Speak() {
	println("peep")
}

type Snake struct {
}

func (s Snake) Eat() {
	println("mice")
}

func (s Snake) Move() {
	println("slither")
}

func (s Snake) Speak() {
	println("hsss")
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	data := make(map[string]Animal)

	for {
		fmt.Print(">")
		scanner.Scan()
		iStr := scanner.Text()
		request := strings.Split(iStr, " ")

		if len(request) < 3 {
			continue
		}
		//type of request
		switch request[0] {
		case "newanimal":
			var a Animal
			//type of animal
			switch request[2] {
			case "cow":
				var t Cow
				a = t
			case "bird":
				var t Bird
				a = t
			case "snake":
				var t Snake
				a = t
			default:
				fmt.Println("Animal not valid. Try again.")
				continue
			}
			data[request[1]] = a
			fmt.Println("Created it!")
		case "query":
			a, b := data[request[1]]

			if b == true {
				switch request[2] {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				default:
					fmt.Println("Query not valid. Try again.")
					continue
				}
			}
		default:
			fmt.Println("Command not valid. Try again.")
			continue
		}

	}
}
