package main

import "fmt"

func main() {
	var status bool = false
	var number int = -56
	name := "Juan"

	day := "lunes"

	if !status {
		// código a ejecutar si la condición es verdadera
		fmt.Println("status =", status)
	}

	if name == "Mario" {
		// código a ejecutar si la condición es verdadera
		fmt.Println("Usuario", name)
	} else if name == "Juan" {
		fmt.Println("Usuario", name)
	}

	if number > 0 {
		fmt.Println("El numero es mayor a 0")
		// código a ejecutar si la condición es verdadera
	} else {
		fmt.Println("El numero es menor a 0")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	nombres := []string{"Juan", "María", "Pedro"}

	for indice, nombre := range nombres {
		fmt.Printf("El nombre en la posición %d es %s\n", indice, nombre)
	}

	switch day {
	case "lunes":
		fmt.Println("Hoy es lunes")
	case "martes":
		fmt.Println("Hoy es martes")
	case "miércoles":
		fmt.Println("Hoy es miércoles")
	case "jueves":
		fmt.Println("Hoy es jueves")
	case "viernes":
		fmt.Println("Hoy es viernes")
	case "sábado":
		fmt.Println("Hoy es sábado")
	case "domingo":
		fmt.Println("Hoy es domingo")
	default:
		fmt.Println("Día inválido")
	}

}
