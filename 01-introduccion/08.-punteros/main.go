package main

import "fmt"

func main() {
	x := 5
	ptr1 := &x          // Puntero a la dirección de memoria de x
	ptr2 := &ptr1       // Puntero a la dirección de memoria de x
	fmt.Println(*ptr1)  // Imprime el valor de x a través del puntero
	fmt.Println(**ptr2) // Imprime el valor de x a través del doble puntero

	modifyValue(&x) // Pasa una referencia a x a través del puntero
	fmt.Println(x)  // Imprime el valor modificado de x

	arreglo := [5]int{1, 2, 3, 4, 5}
	ptr := &arreglo[0]

	for i := 0; i < len(arreglo); i++ {
		fmt.Printf("El valor del elemento %d es %d\n", i, *ptr)
	}

}

func modifyValue(ptr1 *int) {
	*ptr1 = 10 // Modifica el valor a través del puntero
}
