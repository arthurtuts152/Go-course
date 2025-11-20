package main

import "fmt"

func main() {
fmt.Println(fatorial(5))
}

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}
func sumofDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumofDigits((n/10))
}
