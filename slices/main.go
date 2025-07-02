package main
import (
	"fmt")

func main() {

	var slice1 []int
	fmt.Printf("size: %d, value %v\n", len(slice1), slice1)

	slice1 = append(slice1, 10, 20, 30, 40, 50)
	fmt.Printf("size: %d, value %v\n", len(slice1), slice1)

	//Practica
	var slice2 []string
	slice2 = append(slice2, "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo")
	fmt.Printf("size: %d, value %v\n", len(slice2), slice2)

	// IN CASE OF ARRAY
	var myArrayVar [6]int
	fmt.Println(myArrayVar) // Imprime: [0 0 0 0 0 0]
	myArrayVar[1] = 2
	myArrayVar[2] = 5
	myArrayVar[3] = 9
	fmt.Println(myArrayVar) // Imprime: [0 2 5 9 0 0]
	mySlice := myArrayVar[2:4] // Slice from index 2 to 4 (exclusive)
	fmt.Printf("Slice from array: %v\n", mySlice) // Imprime: [5 9]
	fmt.Printf("Size: %d\n", len(mySlice))

	fmt.Println(&mySlice[0])
	fmt.Println(&mySlice[1])

	fmt.Println(myArrayVar)

	mySlice[0] = 80
	mySlice[1] = 90
	fmt.Println(myArrayVar) // Imprime: [0 2 80 90


	
	fmt.Println(myArrayVar[:4])
	fmt.Println(myArrayVar[2:]) // Imprime: [80 90]

	slice := make([]int, 3) // Crea un slice de enteros con longitud 3 y capacidad 3
	fmt.Println(slice) // Imprime: [0 0 0]
	fmt.Printf("Size: %d, value %v\n", len(slice), slice)




}