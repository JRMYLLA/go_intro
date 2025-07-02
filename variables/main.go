package main

import (
	
	"fmt"
	"strconv"
	"unsafe"
)

func main(){
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	const myFirstConst = "a12"
	fmt.Println("My constant is: ", myFirstConst)

	const myIntConst int = 12
	fmt.Println("My integer constant is: ", myIntConst)
	const myStringConst string = "Hello, World!"

	fmt.Println("My string constant is: ", myStringConst)

	
	/*
	int 8
	int 16
	int 32
	int 64
	*/
	fmt.Println()
	var my8BitIntVar int8
	fmt.Printf("Int default value: %d\n", my8BitIntVar)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitIntVar,unsafe.Sizeof(my8BitIntVar),unsafe.Sizeof(my8BitIntVar)*8)

	var my16bitsIntVar int16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16bitsIntVar, unsafe.Sizeof(my16bitsIntVar), unsafe.Sizeof(my16bitsIntVar)*8)
	var my32bitsIntVar int32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32bitsIntVar, unsafe.Sizeof(my32bitsIntVar), unsafe.Sizeof(my32bitsIntVar)*8)
	var my64bitsIntVar int64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n",	 my64bitsIntVar, unsafe.Sizeof(my64bitsIntVar), unsafe.Sizeof(my64bitsIntVar)*8)

	//ahora los uint8, uint16, uint32, uint64
	var my8BitUintVar uint8
	fmt.Printf("type: %T, bytes: %d, bits: %d\n",	 my8BitUintVar, unsafe.Sizeof(my8BitUintVar), unsafe.Sizeof(my8BitUintVar)*8)
	var my16bitsUintVar uint16		
	fmt.Printf("type: %T, bytes: %d, bits: %d\n",	 my16bitsUintVar, unsafe.Sizeof(my16bitsUintVar), unsafe.Sizeof(my16bitsUintVar)*8)
	var my32bitsUintVar uint32		
	fmt.Printf("type: %T, bytes: %d, bits: %d\n",	 my32bitsUintVar, unsafe.Sizeof(my32bitsUintVar), unsafe.Sizeof(my32bitsUintVar)*8)
	var my64bitsUintVar uint64	
	fmt.Printf("type: %T, bytes: %d, bits: %d\n",	 my64bitsUintVar, unsafe.Sizeof(my64bitsUintVar), unsafe.Sizeof(my64bitsUintVar)*8)	
	fmt.Println("End of the program")
	fmt.Println("End of the program")

	//Datos flotantes
	var myFloat32Var float32
	fmt.Printf("Float default value: %f\n", myFloat32Var)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)
	var myfloat64Var float64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myfloat64Var, unsafe.Sizeof(myfloat64Var), unsafe.Sizeof(myfloat64Var)*8)	

	fmt.Println("-------------------------------")
	var myStringVar3 string
	fmt.Printf("String default value: %s\n", myStringVar3)
	//ahora vamos a ver como hacer un string multilinea
	myStringVar5 := "My string variable in golang\nwith multiple \nlines."
	fmt.Println(myStringVar5)

	//Conversiones de variables
	{
		fmt.Println()
		floatVar:= 3.14
		fmt.Printf("type:%T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type:%T, value: %s\n", floatStrVar, floatStrVar)
		
		intVar := 42
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)
		
		intVal1, err := strconv.ParseInt("1333", 0,64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)

		//FLOTANTES
		floatVar1, _ := strconv.ParseFloat("-11.23", 64)
		fmt.Printf("type: %T, value: %f\n", floatVar1, floatVar) 
	

		//BIT: es UN DÍGITO, es el valor más pequeño que puede 
		// tener una variable, es un 0 o un 1
		// BYTE: es un conjunto de 8 bits, es el valor más 
		// pequeño que puede 0=00000000, 1=00000001, 2=00000010, 3=00000011, etc.
		// 255=11111111, 256=0000000100000000
		//KILOBYTE: es un conjunto de 1024 bytes, es el valor más 
		// pequeño que puede tener una variable, es un 0 o un 1
		//MEGABYTE: es un conjunto de 1024 kilobytes,
		
	}

}

