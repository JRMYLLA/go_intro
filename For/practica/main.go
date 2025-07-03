
// Exercise: the following code defines an array of integers and uses a for loop to add 20 to each element of the array. Print the modified array at the end.
package main
import (
	"fmt")
func main() {

	array := [5]int{4,2,5,6,7}

	for i := 0; i <len(array); i++ {
		array[i] = array[i] + 20
	}

	fmt.Println("Los valores del array son: ", array)


}