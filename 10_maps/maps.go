package main

import (
	"fmt"
	"maps"
)
func main() {
	m := make(map[string]string)
	m["name"] = "Golang"
	m["type"] = "Programming Language"
	m["origin"] = "Google"

	fmt.Println("Map:", m)

	// Accessing a value
	fmt.Println("Name:", m["name"])

	// deleting a key-value pair
	delete(m,"type")
	fmt.Println("After deletion:", m)

	// Attempting to access a deleted key
	fmt.Println("Type:", m["type"])//enmpty string
	// clear(m) //empty map

	//checking element in map

	v, ok := m["origin"]

	if ok {
		fmt.Println("Origin exists with value:", v, ok)
	}else{
		fmt.Println("Origin key does not exist",v, ok)
	}

	//eqality of maps

	m2 := map[string]string{"a":"hulhuh", "b":"qweqwe"}

	fmt.Println("maps equal ",maps.Equal(m,m2))
}