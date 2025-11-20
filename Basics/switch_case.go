package main
import "fmt"
func main(){
	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("Its an apple.")
	case "banana":
		fmt.Println("Its a banana.")
	default:
		fmt.Println("Unknown Fruit!")
	}

	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Its a weekday")
	case "Saturday", "Sunday":
		fmt.Println("Its a weekend")
	default:
		fmt.Println("Invalid day")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

// usando fallthrough, quando um caso é encontrado e atende a condição, ele analisa os próximos casos, por default ao achar o primeiro caso que obedece a condição, ele para.

}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Its an integer")
	case float64:
		fmt.Println("Its float")
	case string:
		fmt.Println("Its a string")
	default:
		fmt.Println("Unknown type")
	}

}