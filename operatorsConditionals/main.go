package main

import (
	"fmt"
)
func practica() {

	//Operators and conditionals in Go 

	yearsOld := 32
	fmt.Println("Operators")
	fmt.Println(yearsOld>30) //true
	fmt.Println(yearsOld<32) //false
	fmt.Println(yearsOld==32) //true
	fmt.Println(yearsOld!=32) //false
	fmt.Println(yearsOld>=32) //true
	fmt.Println(yearsOld<=32) //true
	fmt.Println(yearsOld>30 && yearsOld<40) //true

	fmt.Println("------------------------")

	//OR the conditionals in go

	fmt.Println(yearsOld<32 || yearsOld==32) //is (false or true) -> true
	fmt.Println(yearsOld<32 || yearsOld==33) //is (false or false) -> false
	fmt.Println(yearsOld<40 || yearsOld ==33) //is (true or true) -> true

	//AND the operation is true only if both conditions are true
	fmt.Println(yearsOld<32 && yearsOld==32) //is (false and true) -> false
	fmt.Println(yearsOld<32 && yearsOld==33) //is (false and false) -> false
	fmt.Println(yearsOld<40 && yearsOld ==32) //is (true and true) -> true


	fmt.Println("------------------------")
	//NOT the operation in go: the NOT operator negates the value of the condition
	fmt.Println(true) //true
	fmt.Println(!true) //false

	fmt.Println(yearsOld<40) //true
	fmt.Println(!(yearsOld<40)) //false

	fmt.Println(yearsOld<25 && yearsOld==32|| yearsOld <40)// (false and true) or (true) -> true
	fmt.Println(yearsOld<25 && yearsOld==32|| yearsOld <30)// (false and true) or (false) -> false


	fmt.Println("------------------------")
	//Condicionales in Go
	yearsOld1 := 20
	if yearsOld1 > 28 {
		fmt.Printf("%d is higher than 18\n", yearsOld1)
	}

	boolVar := true
	if boolVar {
		fmt.Println("Is true")
	} else {
		fmt.Println("Is false")
	}

	if value :=true; value == true{
		fmt.Println("is true")
	}
	// various conditionals
	number := 10
	if number == 1 {
		fmt.Println("one")
	} else if number == 2 {
		fmt.Println("two")
	}else if number == 3 {
		fmt.Println("three")
	} else {
		fmt.Println("other number")
	}

	//other conditional is switch
	switch number {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		default:
			fmt.Println("other number")




}
// switch with multiple cases

switch{
case number > 0:
	fmt.Println("positive number")
case number < 0:
	fmt.Println("negative number")
	case number == 0:
	fmt.Println("zero")
}
}