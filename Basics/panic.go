package main

import "fmt"

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before panic")
		panic("Input must be a non-negative number")
	}
	fmt.Println("Processing input:", input)
}
func main() {
	//panic(interface{}) - tudo depois de panic não é executado
	process(10)
}