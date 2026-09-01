package main

import (
	"fmt"
)

func main() {
	var A int

	fmt.Scan(&A)

	// list := []int{}

	result := 0

	for divider := 2; divider*divider <= A; divider++ {
		for A%divider == 0 {
			A = A / divider
			// list = append(list, divider)
			result++
		}
	}

	if A > 1 {
		result++
	}

	// fmt.Println(list)
	// fmt.Println(len(list))
	fmt.Println(result)
}
