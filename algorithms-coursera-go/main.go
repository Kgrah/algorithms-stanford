package main

import "fmt"

func main() {
	nums := []int{6, 5, 4, 3, 2, 1}

	inversionsCount := countInversion(nums)
	fmt.Println(inversionsCount)
}
