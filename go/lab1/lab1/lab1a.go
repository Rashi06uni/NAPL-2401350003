package main


import "fmt"

func main() {

// Declare two integer variables
var a int = 20
var b int = 10

// Declare two floating-point variables
var x float64 = 15.5
var y float64 = 5.5

fmt.Println("===== INTEGER OPERATIONS =====")

fmt.Println("First Integer:", a)
fmt.Println("Second Integer:", b)

fmt.Println("Addition:", a+b)
fmt.Println("Subtraction:", a-b)
fmt.Println("Multiplication:", a*b)

fmt.Println("\n===== FLOAT OPERATIONS =====")

fmt.Println("First Float:", x)
fmt.Println("Second Float:", y)

fmt.Println("Addition:", x+y)
fmt.Println("Subtraction:", x-y)
fmt.Println("Multiplication:", x*y)

fmt.Println("\nCalculator Program Completed!")
}