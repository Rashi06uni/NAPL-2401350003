
package main

import "fmt"

func square(ch chan string) {
	fmt.Println("Square goroutine started")
	ch <- fmt.Sprintf("Square of 5 = %d", 5*5)
	fmt.Println("Square goroutine completed")
}

func cube(ch chan string) {
	fmt.Println("Cube goroutine started")
	ch <- fmt.Sprintf("Cube of 3 = %d", 3*3*3)
	fmt.Println("Cube goroutine completed")
}

func fibonacci(ch chan string) {
	fmt.Println("Fibonacci goroutine started")
	a, b := 0, 1
	for i := 0; i < 7; i++ {
		ch <- fmt.Sprintf("%d", a)
		a, b = b, a+b
	}
	fmt.Println("Fibonacci goroutine completed")
}

func main() {
	ch := make(chan string, 10)

	go square(ch)
	go cube(ch)
	go fibonacci(ch)

	for i := 0; i < 9; i++ {
		fmt.Println(<-ch)
	}
}
