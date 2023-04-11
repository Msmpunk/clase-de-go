package main

import "fmt"

func main() {
	resultado := suma(5, 6)
	resultado2 := sum(1, 3)
	resultadoEsPrimo := esPrimo(2)
	numeros := [5]int{1, 20, 50, 4, 5}
	fmt.Println("El valor mas grande en el arreglo es: ", max(numeros[:]))
	fmt.Println(resultado)
	fmt.Println(resultado2)
	fmt.Println(resultadoEsPrimo)

	cociente, err := dividir(10, 0)
	if err != nil {
		// Manejar el error
		fmt.Println("No se puede dividir entre 0")
	} else {
		// Usar el valor del cociente
		fmt.Println("El valor es:", cociente)
	}

}

func suma(a int, b int) int {
	return a + b
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func esPrimo(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func max(nums []int) int {
	if len(nums) == 0 {
		panic("slice vacío")
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		// Retornamos un error si el divisor es cero
		return 0, fmt.Errorf("El divisor no puede ser cero")
	}
	// Retornamos el cociente y un valor nulo para el error
	return a / b, nil
}
