package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (a *Animal) InitMe(n, f, l, no string) {
	a.name = n
	a.food = f
	a.locomotion = l
	a.noise = no
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func SearchAnimal(d []Animal, aName string) Animal {
	var tempA Animal
	tempA.InitMe("", "", "", "")

	for i := range d {
		if d[i].name == aName {
			// Found!
			tempA = d[i]
		}
	}

	return tempA
}

func main() {
	data := make([]Animal, 0, 3)

	var a1 Animal
	a1.InitMe("cow", "grass", "walk", "moo")
	data = append(data, a1)
	a1.InitMe("bird", "worms", "fly", "peep")
	data = append(data, a1)
	a1.InitMe("snake", "mice", "slither", "hsss")
	data = append(data, a1)

	fmt.Println("Valid data:", data)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		scanner.Scan()
		iStr := scanner.Text()
		request := strings.Split(iStr, " ")

		if len(request) < 2 {
			continue
		}

		//Search for the Animal in data
		a2 := SearchAnimal(data, request[0])

		if a2.food != "" {
			switch request[1] {
			case "eat":
				fmt.Println("answer: ", a2.Eat())
			case "move":
				fmt.Println("answer: ", a2.Move())
			case "speak":
				fmt.Println("answer: ", a2.Speak())
			default:
				fmt.Println("try with a valid information!")
			}
		}
	}
}
