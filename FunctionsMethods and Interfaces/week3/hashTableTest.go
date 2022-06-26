package main

import "fmt"

func main(){
 
	idMap := map[string]int { "joe":123}
 
	//adding a key/value pair
	idMap["jane"] = 234
	idMap["alma"] = 345
 
	//deleting key/value pair
	delete(idMap, "joe")

	//Return zero if key is not present
	fmt.Println(idMap["joe"])

	for key, val := range idMap {
		fmt.Println(key, val)
	}

	//two-value assigment tests for existence of the key

	id, p := idMap["joe"]

	fmt.Printf("id: %d, p: %t \n", id, p)
}