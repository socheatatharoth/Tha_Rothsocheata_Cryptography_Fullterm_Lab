package main
import "fmt"
func main(){
	
	var num1, num2 int

	fmt.Print("Plese enter the first number:")
	fmt.Scanln(&num1)

	fmt.Print("Enter the second number:")
	fmt.Scanln(&num2)

	bothpositive := num1 > 0 && num2 > 0
	fmt.Println("Both numbers are positive", bothpositive)

	Onegreater := num1 > num2 || num2 > num1
	fmt.Println("Atleast one number is greater than", Onegreater)

	Notequal := num1 != num2
	fmt.Println("Both number are not equal", Notequal)
}
