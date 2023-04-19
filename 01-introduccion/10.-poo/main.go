package main

import "fmt"

type Direccion struct {
    calle  string
    ciudad string
    estado string
}

type Persona struct {
    nombre    string
    edad      int
    direccion Direccion
}

type Saludador interface {
    saludar()
}

type Animal interface {
    Sonido() string
}

type Perro struct {
    nombre string
}



func main() {
	var s Saludador
	s = Persona{"Juan", 25, Direccion{"Calle 123", "Ciudad de México", "CDMX"}}
	s.saludar()

	p := Persona{"Juan", 12, Direccion{"Calle 123", "Ciudad de México", "CDMX"}}
	fmt.Printf("La dirección de %s es %s, %s, %s\n", p.nombre, p.direccion.calle, p.direccion.ciudad, p.direccion.estado)

	p2 := Persona{"Juan", 30, Direccion{"Calle 123", "Ciudad de México", "CDMX"}}
	p2.presentarse()

	var miAnimal Animal
	miAnimal = Perro{"Fido"}
	fmt.Println(miAnimal.Sonido()) // Imprime "Guau!"
}

func (p Persona) presentarse() {
    fmt.Printf("Hola, mi nombre es %s y tengo %d años\n", p.nombre, p.edad)
}

func (p Persona) saludar() {
    fmt.Printf("Hola, mi nombre es %s\n", p.nombre)
}
func (p Perro) Sonido() string {
    return "Guau!"
}