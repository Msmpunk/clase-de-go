package main

import "fmt"

func main() {
	var nombre string = "Mario"
	name := "Juan"

	var edad int = 25
	age := 30

	var resultSum int = suma(4, 5)
	result := suma(4, 5)

	var a int
	var b float64
	var c string
	var d bool

	const pi = 3.141592

	// Declaración y asignación de variables
	var x float32 = 3.14159265358979323846
	var y float64 = 3.14159265358979323846

	// Declaración y asignación de variables
	var z1 complex64 = complex(1.2, 3.4)
	var z2 complex64 = complex(5.6, 7.8)

	sum := z1 + z2

	// Declaración y asignación de variables
	var r1 rune = 'a'
	var r2 rune = '世'

	// Declaración y asignación de variables
	var u1 uint = 42
	var u2 uint8 = 255
	var u3 uint16 = 65535
	var u4 uint32 = 4294967295
	var u5 uint64 = 18446744073709551615
	var u6 uintptr = 0x7fffffff

	// Imprimir los valores de las variables
	fmt.Println("u1 =", u1)
	fmt.Println("u2 =", u2)
	fmt.Println("u3 =", u3)
	fmt.Println("u4 =", u4)
	fmt.Println("u5 =", u5)
	fmt.Println("u6 =", u6)

	// Imprimir los valores de las variables
	fmt.Println("x =", x)
	fmt.Println("y =", y)

	// Imprimir los valores de las variables
	fmt.Println("z1 =", z1)
	fmt.Println("z2 =", z2)

	fmt.Println("z1 + z2 =", sum)

	// Imprimir los valores de las variables
	fmt.Println("r1 =", r1)
	fmt.Println("r2 =", r2)

	// Imprimir los caracteres correspondientes a los valores de las variables
	fmt.Println("Carácter correspondiente a r1 =", string(r1))
	fmt.Println("Carácter correspondiente a r2 =", string(r2))

	fmt.Println(name, age)
	fmt.Println(nombre, edad)
	fmt.Println(suma(4, 5))
	fmt.Println(resultSum)
	fmt.Println(result)
	fmt.Println(a, b, c, d)
	fmt.Printf("pi: %v\n", pi)
}

func suma(a int, b int) int {
	return a + b
}
