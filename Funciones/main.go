// Functions in Go are defined using the `func` keyword, followed by the function name, parameters, and return type. Here's a simple example of a function that adds two integers:
package main

import (
	"fmt"

	"github.com/JRMYLLA/go_intro/Funciones/function"
)

func main() {
	fmt.Println(function.Add(3, 4))
	
	function.RepeatString(20, "saz")
	fmt.Println() // New line for better readability

	function.RepeatString(2, "Go")

	fmt.Println()
	// Using the Calc function from the function package

	v, err := function.Calc(function.SUM, 3,6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
	fmt.Println("Value: ", v)
	}
	v, err = function.Calc(function.DIV, 3,0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Value: ", v)
	}

	x, y := function.Split(20)
	fmt.Println("Value 1: ", x, "Value 2: ", y)
	fmt.Println()

	v=function.MSum(23,12,32,12,3,1,2,3,2,1,23,12,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20)
	fmt.Println("multy sum: ", v)

	fmt.Println("------------------------------")


	v, err = function.MOperations(function.SUM, 1,2,3,4,5,6,7,8,9,10)
	fmt.Println("multy sum: ", v, "Error: ", err)
	v, err = function.MOperations(function.DIV, 1,2,3,4,5,6,7,8,9,10)
	fmt.Println("multy div: ", v, "Error: ", err)
	v, err = function.MOperations(function.MUL, 1,2,3,4,5,6,7,8,9,10)
	fmt.Println("multy mul: ", v, "Error: ", err)
	v, err = function.MOperations(function.SUB, 1,2,3,4)
	fmt.Println("multy sub: ", v, "Error: ", err)
	fmt.Println("End of the program")


}
