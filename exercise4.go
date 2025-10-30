package main

import "fmt"

// Function to add two numbers
func add(num1, num2 float64) float64 {
	return num1 + num2
}

// Function to subtract two numbers
func sub(num1, num2 float64) float64 {
	return num1 - num2
}

// Function to multiply two numbers
func mul(num1, num2 float64) float64 {
	return num1 * num2
}

// Function to divide two numbers
func div(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return num1 / num2, nil
}

// Function to find modulus of two numbers
func mod(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return num1 % num2, nil
}

func main() {
	var choice int

	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit")
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting...")
			break
		}

		var num1, num2 float64
		var aInt, bInt int
		var result float64
		var intResult int
		var err error

		fmt.Print("Enter num1: ")
		fmt.Scan(&num1)
		fmt.Print("Enter num2: ")
		fmt.Scan(&num2)

		// Convert to int for modulus operation
		aInt = int(num1)
		bInt = int(num2)

		switch choice {
		case 1:
			result = add(num1, num2)
			fmt.Println("Result:", result)
		case 2:
			result = sub(num1, num2)
			fmt.Println("Result:", result)
		case 3:
			result = mul(num1, num2)
			fmt.Println("Result:", result)
		case 4:
			result, err = div(num1, num2)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Result:", result)
			}
		case 5:
			intResult, err = mod(aInt, bInt)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Result:", intResult)
			}
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
