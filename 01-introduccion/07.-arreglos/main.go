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
	persona3 := Persona{
		nombre: "pedro",
		edad:   25,
	}

	arreglo := [3]Persona{
		persona1,
		persona2,
		persona3,
	}

	personas := make([]Persona, 0)
	personas = append(personas, persona3)
	personas = append(personas, persona1)

	for i := 0; i < len(personas); i++ {
		fmt.Println("Nombre:", personas[i].nombre, "- Edad:", personas[i].edad)
	}

	for i := 0; i < len(arreglo); i++ {
		fmt.Println("Nombre:", arreglo[i].nombre, "- Edad:", arreglo[i].edad)
	}

	// Crear un mapa con valores iniciales
	notas := map[string]float64{
		"Juan":  8.5,
		"María": 9.0,
		"Pedro": 7.5,
	}

	// Iterar sobre el mapa y imprimir los valores
	for name, num := range notas {
		fmt.Printf("%s: %.2f\n", name, num)
	}

}
