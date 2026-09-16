package main

import (
	"fmt"
	"math"
)

func main() {
	var input int

	fmt.Scan(&input)

	var binary []int

	for input >= 1 {
		binary = append(binary, input%2)
		input /= 2
	}

	var output float64
	binaryLen := len(binary)

	for i := 0; i < binaryLen; i++ {
		output += math.Pow(2, float64(binaryLen)-float64(i)-1) * float64(binary[i])
	}

	fmt.Println(int(output))
}
