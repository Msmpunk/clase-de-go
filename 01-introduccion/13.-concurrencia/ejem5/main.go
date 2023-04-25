package main

import "fmt"

func sumaConcurrente(nums []int) int {
    n := len(nums)
    c := make(chan int)

    for i := 0; i < n; i++ {
        go func(start int) {
            sum := 0
            for j := start; j < n; j += len(nums) {
                sum += nums[j]
            }
            c <- sum
        }(i)
    }

    total := 0
    for i := 0; i < n; i++ {
        total += <-c
    }

    return total
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    resultado := sumaConcurrente(nums)
    fmt.Println("La suma es:", resultado)
}
