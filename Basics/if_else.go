package main
import "fmt"

func main() {
	age := 25

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	temperature := 25
	if temperature >= 30 {
		fmt.Println("Its hot outside.")
	} else {
		fmt.Println("Its cold outside")
	}

	score := 85

	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}
	// Se eu tirasse o else e deixasse apenas o if, todas as condições satisfeitas printariam

	num := 15

	if num % 2 == 0 {
		if num % 3 == 0 {
			fmt.Println("Number is divisible by both 2 and 3")
		} else {
			fmt.Println("Number is divisible by 2 but not 3")
		}
	} else {
		fmt.Println("Number is not divisible by 2")
	}

	if 10%2 == 0 || 5 % 2 == 0 {
		fmt.Println("Either 10 or 5 are even")
	}
	if 10%2 == 0 && 6%2 == 0 {
		fmt.Println("Both 10 and 6 are even")
	} 
}