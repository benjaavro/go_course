package main

import (
	"fmt"
	"maps"
)

func main() {
	// var name map[keyType]valueType

	// mapVariable := make(map[keyType]valueType)
	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["one"] = 1
	myMap["two"] = 2
	fmt.Println(myMap)
	fmt.Println(myMap["one"])
	fmt.Println(myMap["three"])
	delete(myMap, "two")
	fmt.Println(myMap)
	myMap["three"] = 3
	clear(myMap)
	fmt.Println(myMap)

	myMap["one"] = 1
	_, ok := myMap["four"]

	if ok {
		fmt.Println("Exists")
	} else {
		fmt.Println("Doesn't exist")
	}

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	myMap3 := map[string]int{"a": 1, "b": 2, "c": 3}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("Maps are the same")
	}

	for _, value := range myMap2 {
		fmt.Println(value)
	}

	// Nil init map
	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("Map is init to nil")
	}

	val, ok2 := myMap4["wii"]
	fmt.Println(val, ok2)
	// myMap["wii"] = "wii" (( INVALID SET ))
	myMap4 = make(map[string]string)
	myMap4["a"] = "A"
	fmt.Println(myMap4)

	// Get Map length
	fmt.Println(len(myMap2))

	// Nested maps
	myBigMap := make(map[string]map[string]string)
	myBigMap["dog"] = make(map[string]string)
	myBigMap["dog"]["belgian"] = "Zuri"
	myBigMap["dog"]["gsp"] = "Nia"
	fmt.Println(myBigMap)
}
