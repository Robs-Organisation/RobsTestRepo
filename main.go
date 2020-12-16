package main

import "fmt"

func add(varOne int, varTwo int) int {
	return varOne + varTwo
}

func main() {

	numberOne := 1
	numberTwo := 2

	result := add(numberOne, numberTwo)

	fmt.Println("The result: ", result)

}
