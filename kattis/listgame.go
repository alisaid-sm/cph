package main

import (
	"fmt"
)

func main() {
	var A int

	fmt.Scan(&A)

	//list := []int{2}

	result := 1

	divider := 2

	for A/divider != 1 {
		if A%divider != 0 {
			divider++
			continue
		}
		A = A / divider
		//list = append(list, divider)
		result++
	}

	//fmt.Println(list)
	//fmt.Println(len(list))
	fmt.Println(result)
}
