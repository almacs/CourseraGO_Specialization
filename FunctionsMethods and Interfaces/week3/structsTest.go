package main

import "fmt"

type Person struct {
	name string
	addr string
	phone string
}


func main(){
	//Person Struct
	//Option1 have 3 separate varaiable, 
	//programmer remembers that they are related


	//Option2 make a single struct which contains all 3 vars
	var p1 Person

	p1.name = "Joe"
	p1.addr = "444 Saratoga Ave."

	x := p1.addr

	fmt.Println(x)


}