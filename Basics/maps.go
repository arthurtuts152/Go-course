package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)
	// using a Map Literal
	// mapVariable = map[keyType]valueType {key1: value1, key2: value2}
	myMap := make(map[string]int)
	myMap["key1"] = 9
	fmt.Println(myMap)
	delete(myMap, "key1")
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)

	//clear(myMap)
	//fmt.Println(myMap)
	value, unknownvalue := myMap["key1"]
	fmt.Println(value)
	fmt.Println(unknownvalue)

	myMap2 := map[string]int{"a" : 1, "b" : 2}
	fmt.Println(myMap2)
	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	}
	myMap3 := map[string]int{"a" : 1, "b" : 2}
	if maps.Equal(myMap2, myMap3) {
		fmt.Println("myMap2 and myMap3 are equal")
	}
	for _, v := range myMap3 {
		fmt.Println(v)
	}
	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is not initialized to nil value")
	}
	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println(myMap4)
}