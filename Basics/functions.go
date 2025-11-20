package main

import "fmt"

func main() {
	// func <name> (parameters) returnType { code block, return value }
	operation := add

	result := operation(3, 5)
	fmt.Println(result)
	x := applyOperation(1,2, operation)
	fmt.Println(x)
	multiplyBy2 := createMultiplier(2)
	m := multiplyBy2(6)
	fmt.Println(m)

}
func add(a int, b int) int {
	return a + b
}
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x,y)
}
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
