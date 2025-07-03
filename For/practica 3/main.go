package main
import (
	
	"fmt"

)

const notebook string = "notebook"
const tv string = "tv"
const heladera string = "heladera"
const monitor string = "monitor"
const camara string = "camara"
const notFounf string = "Producto no encontrado"

func main() {

	var array []string
	fmt.Println("Selecciones valores, sale con '0'")

	for {

		var value string
		fmt. Scanf("%s", &value)
		if value == "0" {
			break
		}
		switch value {
		case "10":
			array = append(array, notebook)
			case "15":
			array = append(array, tv)
			case "21":
				array = append(array, heladera)
			case "27":
				array = append(array, monitor)
			case"35":
				array = append(array, camara)
		default:
			array = append(array, notFounf)
	}





}
fmt.Printf("Los valores del array son: %v\n", array)
}