package main

import (
	"fmt"
)
func main() {

	m1 := make(map[int]string)
	m1[1] = "A"
	m1[2] = "B"
	m1[3] = "C"

	fmt.Println(m1)
	fmt.Println(m1[1])

	delete(m1, 2) // Elimina la clave 2
	fmt.Println(m1)

	m1[1] ="A2"
	fmt.Println(m1[1])
	m1[7] = ""
	fmt.Println(m1[7]) // Imprime: "" (cadena vacía)
	fmt.Println(m1[99]) // Imprime: "" (cadena vacía, clave no existe)

	v, ok := m1[7]
	fmt.Println("The value is:", v, "and the key exists?", ok) // Imprime: The value is:  and the key exists? true

	v, ok =m1[99]
	fmt.Println("The value:", v, "Present?", ok) // Imprime: The value:  Present? false

	m2 := map[int]string{
		4 : "A",
		5 : "C",
		6 : "D",
	
}
	fmt.Print(m2) // Imprime: map[4:A 5:C 6:D]

	
}