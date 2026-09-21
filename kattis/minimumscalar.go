package main

import (
	"fmt"
	"sort"
)

/*
1 3 -5 -> -5 1 3 -> -20 1 -6 -> -25
-2 4 1 -> 4 1 -2

1 2 3 4 5 -> 1 2 3 4 5 -> 1 2 3 0 0 -> 6
1 0 1 0 1 -> 1 1 1 0 0
*/

func main() {
	var T int

	fmt.Scan(&T)

	vectors := map[int][][]int{}

	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)

		for j := 0; j < 2; j++ {
			for k := 0; k < n; k++ {
				var v int
				fmt.Scan(&v)

				if vectors[i] == nil {
					vectors[i] = make([][]int, T)
				}

				vectors[i][j] = append(vectors[i][j], v)
			}

			sort.Slice(vectors[i][j], func(a, b int) bool {
				if j == 0 {
					return vectors[i][j][a] < vectors[i][j][b]
				} else {
					return vectors[i][j][a] > vectors[i][j][b]
				}
			})
		}

		var sum int

		for k := 0; k < len(vectors[i][0]); k++ {
			sum += vectors[i][0][k] * vectors[i][1][k]
		}

		fmt.Println(fmt.Sprintf("Case #%d: %d", i+1, sum))
	}
}
