package main
import "fmt"
func main() {
	 var numbers[5]int
	 fmt.Println(numbers) 

	 numbers[4] = 20
	 fmt.Println(numbers)

	 numbers[0] = 9
	 fmt.Println(numbers)

	 fruit := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	 fmt.Println(fruit)

	 fmt.Println("Third fruit:", fruit[2])

	 original:= [3]int{1,2,3}
	 copied := original
	 copied[0] = 100

	 fmt.Println("Original array:", original)
	 fmt.Println("Copied array:", copied)

	 for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
		// Não se pode usar println quando usa-se placeholders, %d, %f, %s, nesse caso usa-se printf
	 }

	 for i, v:= range numbers {
		fmt.Printf("Index: %d, Value %d\n", i, v)
	 }
	 // Blank identifier - caso não se queira identificar uma variável, usa se o _

	 b := 2
	 _ = b

	 fmt.Println("The lenght of numbers array is:", len(numbers))
	 a1 := [3]int{1,2,3}
	 a2 := [3]int{1,2,3}
	 fmt.Println(a1 == a2)

	 var matrix [3][3]int = [3][3]int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	 }
	 fmt.Println(matrix)

}