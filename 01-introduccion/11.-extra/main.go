package main

import (
    "fmt"
    "strings"
)

type Persona struct {
    nombre    string
    direccion string
}

func main() {
    var personas []Persona

    for {
        fmt.Println("1. Registrar persona")
        fmt.Println("2. Eliminar persona")
        fmt.Println("3. Mostrar personas registradas")
        fmt.Println("4. Salir")
        var opcion int
        fmt.Print("Ingrese una opcion: ")
        fmt.Scanln(&opcion)

        switch opcion {
        case 1:
            var nombre, direccion string
            fmt.Print("Ingrese el nombre de la persona: ")
            fmt.Scanln(&nombre)
            fmt.Print("Ingrese la direccion de la persona: ")
            fmt.Scanln(&direccion)
            personas = append(personas, Persona{nombre: strings.TrimSpace(nombre), direccion: strings.TrimSpace(direccion)})
            fmt.Println("Persona registrada correctamente.")
        case 2:
            var indice int
            fmt.Print("Ingrese el indice de la persona a eliminar: ")
            fmt.Scanln(&indice)
            if indice >= 0 && indice < len(personas) {
                personas = append(personas[:indice], personas[indice+1:]...)
                fmt.Println("Persona eliminada correctamente.")
            } else {
                fmt.Println("Indice invalido.")
            }
        case 3:
            fmt.Println("Personas registradas:")
            for i, persona := range personas {
                fmt.Printf("%d. Nombre: %s | Direccion: %s\n", i, persona.nombre, persona.direccion)
            }
        case 4:
            fmt.Println("Hasta luego.")
            return
        default:
            fmt.Println("Opcion invalida.")
        }
    }
}