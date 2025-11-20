package main

import (
	"fmt"
	"errors"
)
func divide(a,b int) (int, int) {
	quotient := a/b
	remainder := a%b
	return quotient, remainder
}

func compare(a,b int) (string, error) {
	if(a > b) {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}
func main() {
	// func <name> (parameter1 type1, parameter2 type2) (returnType1, returnType2) { code block, returnValue1, returnValue2}
	q,r := divide(10,3)
	fmt.Printf("Quotient %d. Remainder %d\n", q, r)

	result, err := compare(3, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}