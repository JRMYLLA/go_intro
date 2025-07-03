package main
import (
	"fmt")

func main() {
	sum := 0

	for i := 0; i<10; i++{
		sum++

	}
	fmt.Println(sum)	
	sum = 1
	for sum<1000{
		sum++
	}
	fmt.Println(sum) // Imprime: 1000

	for{
		if sum > 1000 {
			break
		}
		sum++
	}
	fmt.Println(sum) // Imprime: 1001

	arr := []int{1, 2, 3, 1, 2, 3}
	fmt.Println()
	for i := range arr {
		fmt.Println("Index: ", i, " - Value: ", arr[i])

}
	fmt.Println("------------------------------------------------------")
	for _,v := range arr {
		fmt.Println(" Value: ", v)
	}


	map2 := map[string]float64{
		"A" : 1.22,
		"B" : 2.33,
		"C" : 3,
	}
	for key, value := range map2 {
		fmt.Println("Key: ", key, " - Value: ", value)
	}

	//Recorrer un mapa con range, sirve para iterar sobre las claves y valores de un mapa.
	fmt.Println("------------------------------------------------------")

	map3 := map[string][]int{
		"A":nil,
		"B":{2,34,1,2,4},
		"C":{3,4,5,6,7},
}
	for key, value := range map3 {
		fmt.Println("Key: ", key)
		for _, v := range value {
			fmt.Println("Value: ", v)
		}
	}
	fmt.Println()
}