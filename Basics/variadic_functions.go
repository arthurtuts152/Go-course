package main
import "fmt"
func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
func main() {
	//...Elipsis
	//func functionName(param1 type1, param2 type2, param3 ...type3) returnType {
	// function body
	//}
	sequence, total := sum(1,20,30,40,50,60)
	fmt.Println("Sequence: ", sequence, "Total", total)
	numbers := []int{1,2,3,4,5,9}
	sequence1, total1 := sum(3, numbers...)
	fmt.Println("Sequence: ", sequence1, "Total", total1)
}