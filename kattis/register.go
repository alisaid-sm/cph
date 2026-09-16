package main

import (
	"fmt"
)

/*
capacitys 2 3 5 7 11 13 17 19
registers 0 0 4 6 10 12 16 18

currentVal 2304
multiplier 210
*/

func main() {
	registers := make([]int, 8)
	registersCapacity := []int{2, 3, 5, 7, 11, 13, 17, 19}

	for i := 0; i < 8; i++ {
		var input int

		fmt.Scan(&input)

		registers[i] = input
	}

	maxValue := 9699689
	currentValue := 0
	multiplier := 1

	for i := 0; i < 8; i++ {
		currentValue += registers[i] * multiplier
		multiplier *= registersCapacity[i]
	}

	fmt.Println(maxValue - currentValue)
}
