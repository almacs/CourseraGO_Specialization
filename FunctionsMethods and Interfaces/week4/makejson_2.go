package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Println("please input name :")
	fmt.Scanf("%s", &name)
	fmt.Println("please input address :")
	fmt.Scanf("%s", &address)

	titucion
	ewJson := &teble{
		Name:    name,
		Address: address,
	}

	jsonStu, err := json.Marshal(newJson)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonStu))
}

type teble struct {
	Name    string `json:"name"`
	Address string
}
