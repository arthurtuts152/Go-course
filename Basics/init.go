package main

import "fmt"
func init() {
	fmt.Println("Inside init")
}
func main() {
	//init executes before main, it initializes variables, if there is more than 1 init, they execute in the order they were written
fmt.Println("Inside the main")
}