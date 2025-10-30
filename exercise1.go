package main

import "fmt"

func main() {
    var a, b int

    fmt.Print("Enter first number (a): ")
    fmt.Scanln(&a)
    fmt.Print("Enter second number (b): ")
    fmt.Scanln(&b)

    // Save original values
    num1 := a
    num2 := b

    fmt.Println("\nInitial values:")
    fmt.Println("a =", num1, ", b =", num2)

    // Assignment =
    a = b
    fmt.Println("\na = b → a =", a)

    // Reset a and b to original values
    a, b = num1, num2

    // Add and assign +=
    a += b
    fmt.Println("a += b → a =", a)
    a, b = num1, num2

    // Subtract and assign -=
    a -= b
    fmt.Println("a -= b → a =", a)
    a, b = num1, num2

    // Multiply and assign *=
    a *= b
    fmt.Println("a *= b → a =", a)
    a, b = num1, num2

    // Divide and assign /=
    if b != 0 {
        a /= b
        fmt.Println("a /= b → a =", a)
    } else {
        fmt.Println("Cannot divide by zero")
    }
    a, b = num1, num2

    // Modulus and assign %=
    if b != 0 {
        a %= b
        fmt.Println("a %= b → a =", a)
    } else {
        fmt.Println("Cannot find modulus by zero")
    }
}
