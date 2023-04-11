package main

import "fmt"

type Dirección struct {
	calle  string
	ciudad string
	estado string
}

type Persona struct {
	nombre    string
	edad      int
	dirección Dirección
}

func main() {

	persona1 := Persona{
		nombre: "Juan",
		edad:   25,
		dirección: Dirección{
			calle:  "Av. Independencia",
			ciudad: "Mex",
			estado: "CDMX",
		},
	}

	persona2 := Persona{
		nombre: "Juan",
		edad:   25,
	}

	personaPtr := &persona1

	fmt.Println((*personaPtr).nombre)
	fmt.Println((*personaPtr).edad)

	fmt.Println(persona2.nombre)
	fmt.Println(persona1.dirección.calle)
	fmt.Println(persona1.dirección.ciudad)
	fmt.Println(persona1.dirección.estado)

	fmt.Println(persona1.nombre)
	fmt.Println(persona1.edad)

}
