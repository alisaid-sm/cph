package main

import "fmt"

func main() {
	var A, B, C int

	fmt.Scan(&A, &B)

	grid := make([][]byte, A)

	for i := 0; i < A; i++ {
		var s string
		fmt.Scan(&s)
		grid[i] = []byte(s)
	}

	for i := 0; i < A; i++ {
		for j := 0; j < len(grid[i]); j++ {
			var left, right, top, down byte

			if j == 0 {
				left = '0'
			} else {
				left = grid[i][j-1]
			}

			if j == len(grid[i])-1 {
				right = '0'
			} else {
				right = grid[i][j+1]
			}

			if i == 0 {
				top = '0'
			} else {
				top = grid[i-1][j]
			}

			if i == A-1 {
				down = '0'
			} else {
				down = grid[i+1][j]
			}

			if grid[i][j] == '1' {
				if left == '0' {
					C++
				}
				if right == '0' {
					C++
				}
				if top == '0' {
					C++
				}
				if down == '0' {
					C++
				}
			} else {
				if left == '1' && right == '1' && top == '1' && down == '1' {
					C -= 4
				}
			}
		}
	}

	fmt.Println(C)
}
