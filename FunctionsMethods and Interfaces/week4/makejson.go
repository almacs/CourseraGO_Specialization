package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter name:")
	scanner.Scan()
	nameInp := scanner.Text()

	fmt.Println("Enter address:")
	scanner.Scan()
	addrInp := scanner.Text()

	d := map[string]string{"Name": nameInp, "Address": addrInp}

	barr, err := json.Marshal(d)

	if err != nil {
		panic(err)
	}

	fmt.Println("Json representation:", string(barr))
}
