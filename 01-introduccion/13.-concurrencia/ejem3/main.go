package main

import (
	"fmt"
	"sync"
)

func sum(nums []int, wg *sync.WaitGroup, mu *sync.Mutex, total *int) {
	defer wg.Done()

	sum := 0
	for _, num := range nums {
		sum += num
	}

	// Proteger la variable total con un Mutex
	mu.Lock()
	*total += sum
	mu.Unlock()
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	var mu sync.Mutex
	total := 0

	// Agregar dos goroutines al WaitGroup
	wg.Add(2)

	// Ejecutar las goroutines en paralelo
	go sum(nums[:len(nums)/2], &wg, &mu, &total)
	go sum(nums[len(nums)/2:], &wg, &mu, &total)

	// Esperar a que todas las goroutines hayan terminado
	wg.Wait()

	fmt.Println("Total:", total)
}
