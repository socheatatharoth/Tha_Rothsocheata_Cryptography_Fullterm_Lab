package main
import "fmt"
func Myxor(num1, num2 int){
	result := num1 ^ num2
	fmt.Printf("XOR of %d and %d = %d",num1, num2, result)

}
func Mynot(num1 int){
	result := ^num1
	fmt.Println("Not of", num1, "is", result)

}

func MyOr(num1, num2 int){
	result := num1 | num2 
	fmt.Println("Or of", num1, "and", num2, "is", result)

}

func Myand(num1, num2 int){
	result := num1 & num2
	fmt.Println("And of", num1, "and", num2, "is:", result)


}

func bitwiseShifts(num1 int){
	fmt.Println("Left shift of", num1, "by 1:", num1<<1)
	fmt.Println("Right shift of", num1, "by 1:", num1>>1)
}

func main(){
	var num1, num2 int
	
	fmt.Print("Enter the first number:")
	fmt.Scan(&num1)

	fmt.Print("Plese enter the second the number:")
	fmt.Scan(&num2)


	Myxor(num1, num2)
	Mynot(num1)
	MyOr(num1, num2)
	Myand(num1, num2)
	bitwiseShifts(num1)
	



}