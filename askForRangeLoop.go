package main

import "fmt"

var numberOne, numberTwo int

func printOddEven(small int, big int, remainder int) {
	println("even numbers are")
	for i := 0; i <= (big - small); i++ {
		if Number := (small + i); Number%2 == remainder {
			println(Number)
		}
	}
}

func main() {
	//Insert code for only even here
	fmt.Println("Input two numbers")
	fmt.Println("Input first number")
	fmt.Scan(&numberOne)
	fmt.Println("Input second number")
	fmt.Scan(&numberTwo)

	if numberOne > numberTwo {
		numberOne, numberTwo = numberTwo, numberOne
	}

	printOddEven(numberOne, numberTwo, 0)
	printOddEven(numberOne, numberTwo, 1)

}
