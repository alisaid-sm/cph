package main

import (
	"fmt"
	"strings"
)

type Point struct {
	i int
	j int
}

func main() {
	var A, B, C int

	fmt.Scan(&A, &B)

	A += 2
	B += 2

	grid := make([][]byte, A)

	for i := 0; i < A; i++ {
		var s string
		if i == 0 || i == A-1 {
			s = strings.Repeat("0", B)
			grid[i] = []byte(s)
			continue
		}
		fmt.Scan(&s)
		s = "0" + s + "0"
		grid[i] = []byte(s)
	}

	visited := make([][]bool, A)

	for i := range visited {
		visited[i] = make([]bool, B)
	}

	visited[0][0] = true

	queue := []Point{
		{0, 0},
	}

	directions := []Point{
		{-1, 0}, // top
		{1, 0},  // bottom
		{0, -1}, // left
		{0, 1},  // right
	}

	for head := 0; head < len(queue); head++ {
		current := queue[head]

		for _, direction := range directions {
			ni := current.i + direction.i
			nj := current.j + direction.j

			if ni < 0 || ni == A || nj < 0 || nj == B {
				continue
			}

			if !visited[ni][nj] {
				if grid[ni][nj] == '0' {
					queue = append(queue, Point{ni, nj})
					visited[ni][nj] = true
				} else {
					C++
				}
			}
		}
	}

	fmt.Println(C)
}
