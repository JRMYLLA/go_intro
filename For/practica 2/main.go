// Debemos implementar un programa en go para ir ingresando valorres por consola hastaque se agrega
// cer y finaliza el programa. Todos los valores ingresados por consola deben
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"

)
func main() {

	//crear un slice para almacenar los números
	numeros:= []int{}

	//Inicializar el lector de entrada
	lector := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese los números enteros (0 para finalizar):")

	for{
		//leer la entrada del usuario
		entrada, _:= lector.ReadString('\n')
		//Eliminar espacios y saltos de línea
		valor := strings.TrimSpace(entrada)
		//Convertir la entrada a un entero
		num, err := strconv.Atoi(valor)
		if err != nil {
			fmt.Println("Entrada no válida, por favor ingrese un número entero.")
			continue
		}
		//Comprobar si el número es cero
		if num == 0 {
			break
		}
		//agregar número al slice
		numeros = append(numeros, num)

	}
	fmt.Printf("Los valores del array son: %v\n", numeros)
	


}