package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B float64

	fmt.Scan(&A, &B)

	fmt.Println(math.Ceil(A / float64(math.Sin(B*math.Pi/180))))
}
