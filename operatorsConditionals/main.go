package main

import (
	"fmt"
)
func main() {

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
}