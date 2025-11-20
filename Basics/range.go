package main
import "fmt"

func main() {

	message := "Hello World"

	for i, v := range message {
		fmt.Println(i, v)
	}
	//arrays e slices o range vai na ordem, para maps e channels ele ietra sobre as chaves e em ordem não especifica


}