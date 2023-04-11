package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("¡Bienvenido al programa de cálculo de áreas!")

	for {
		fmt.Println("\nSeleccione una opción:")
		fmt.Println("1. Calcular el área de un triángulo")
		fmt.Println("2. Calcular el área de un pentágono")
		fmt.Println("3. Calcular el área de un círculo")
		fmt.Println("4. Salir del programa")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("\nHa seleccionado calcular el área de un triángulo.")
			var base, altura float64
			fmt.Println("Ingrese la base del triángulo:")
			fmt.Scanln(&base)
			fmt.Println("Ingrese la altura del triángulo:")
			fmt.Scanln(&altura)
			area := (base * altura) / 2
			fmt.Printf("El área del triángulo es: %.2f\n", area)

		case 2:
			fmt.Println("\nHa seleccionado calcular el área de un pentágono.")
			var lado, apotema float64
			fmt.Println("Ingrese el lado del pentágono:")
			fmt.Scanln(&lado)
			fmt.Println("Ingrese la apotema del pentágono:")
			fmt.Scanln(&apotema)
			perimetro := lado * 5
			area := (perimetro * apotema) / 2
			fmt.Printf("El área del pentágono es: %.2f\n", area)

		case 3:
			fmt.Println("\nHa seleccionado calcular el área de un círculo.")
			var radio float64
			fmt.Println("Ingrese el radio del círculo:")
			fmt.Scanln(&radio)
			area := math.Pi * math.Pow(radio, 2)
			fmt.Printf("El área del círculo es: %.2f\n", area)

		case 4:
			fmt.Println("\n¡Gracias por usar el programa!")
			return

		default:
			fmt.Println("\n¡Opción inválida! Seleccione una opción del menú.")
		}
	}
}
