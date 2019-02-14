package main
//import fmt library
import "fmt"
//main function
func main() {
	fmt.Println("Hello go")
	result:= addition(1,2)
	fmt.Println(result)
}
//function to add two nums 
func addition (a int, b int) int {
	sum:= a+b
	return sum
}