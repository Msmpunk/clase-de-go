package main

import (
	"fmt"
)

func sum(nums []int, c chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	c <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	c1 := make(chan int)
	c2 := make(chan int)

	go sum(nums[:len(nums)/2], c1)
	go sum(nums[len(nums)/2:], c2)

	sum1, sum2 := <-c1, <-c2

	fmt.Println("Sum1:", sum1)
	fmt.Println("Sum2:", sum2)
	fmt.Println("Total:", sum1+sum2)
}
