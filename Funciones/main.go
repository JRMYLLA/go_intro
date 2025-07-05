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

}
