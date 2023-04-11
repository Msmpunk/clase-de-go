package main

import "fmt"

type ErrorPersonalizado struct {
	mensaje string
}

func main() {
	// Llamada a la función
	resultado1 := dividir(10, 2)

	fmt.Println("El resultado es:", resultado1)

	// Llamada a la función
	resultado2, err := dividir2(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El valor es:", resultado2)
	}

	// Llamada a la función
	resultado3, err := dividir3(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El valor es:", resultado3)
	}
}

func dividir(a, b float64) float64 {
	if b == 0 {
		panic("El divisor no puede ser cero Error panic")
	}
	return a / b
}

func dividir2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("El divisor no puede ser cero, error tipo 2")
	}
	return a / b, nil
}

func (e *ErrorPersonalizado) Error() string {
	return e.mensaje
}

func dividir3(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &ErrorPersonalizado{"El divisor no puede ser cero, error tipo 3"}
	}
	return a / b, nil
}
