package main
import (
	"fmt"
)
func main() {
	// Practica de operadores y condicionales en Go
	// Dentro de nuestro codigo vamos a tener 
	//2 variables definidas, que van a ser:  licencia y edad

	licence := true
	edad := 15
	
	if edad > 15 && licence {
		fmt.Println("Puede seguir avanzando")
	}else{
		fmt.Println("No puede seguir avanzando")
	}

}