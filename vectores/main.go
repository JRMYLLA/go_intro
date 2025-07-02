package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// The vector is a data structure that allows you to store multiple values of the same type.
	var myArrayVar [6]int
	fmt.Println(myArrayVar)

	myArrayVar1 := [3] string{"a", "b", "c"}
	fmt.Println(myArrayVar1)
	myArrayVar[1]=2
	myArrayVar[2]=5
	myArrayVar[3]=9
	fmt.Println(myArrayVar)

	fmt.Println("Size:", len(myArrayVar))
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)



}