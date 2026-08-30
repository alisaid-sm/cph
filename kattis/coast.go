package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

	visited := make(map[string]struct{})
	queue := []string{}

	var i, j, k int

	for len(queue) > 0 || i+j == 0 {
		k++
		visitedIndex := strconv.Itoa(i) + "," + strconv.Itoa(j)
		// fmt.Println(visitedIndex, k)

		// top
		if i != 0 {
			topIndex := strconv.Itoa(i-1) + "," + strconv.Itoa(j)
			_, exists := visited[topIndex]
			if grid[i-1][j] == '0' && !exists {
				queue = append(queue, topIndex)
				visited[topIndex] = struct{}{}
			} else if grid[i-1][j] == '1' {
				C++
				visited[topIndex] = struct{}{}
			}
		}

		// bottom
		if i != A-1 {
			bottomIndex := strconv.Itoa(i+1) + "," + strconv.Itoa(j)
			_, exists := visited[bottomIndex]
			if grid[i+1][j] == '0' && !exists {
				// fmt.Println("bottomIndex", bottomIndex)
				queue = append(queue, bottomIndex)
				visited[bottomIndex] = struct{}{}
			} else if grid[i+1][j] == '1' {
				C++
				visited[bottomIndex] = struct{}{}
			}
		}

		// right
		if j != B-1 {
			rightIndex := strconv.Itoa(i) + "," + strconv.Itoa(j+1)
			_, exists := visited[rightIndex]
			if grid[i][j+1] == '0' && !exists {
				queue = append(queue, rightIndex)
				visited[rightIndex] = struct{}{}
			} else if grid[i][j+1] == '1' {
				C++
				visited[rightIndex] = struct{}{}
			}
		}

		// left
		if j != 0 {
			leftIndex := strconv.Itoa(i) + "," + strconv.Itoa(j-1)
			_, exists := visited[leftIndex]
			if grid[i][j-1] == '0' && !exists {
				queue = append(queue, leftIndex)
				visited[leftIndex] = struct{}{}
			} else if grid[i][j-1] == '1' {
				C++
				visited[leftIndex] = struct{}{}
			}
		}

		visited[visitedIndex] = struct{}{}

		// fmt.Println(visited)

		dequeue := queue[0]

		next := strings.Split(dequeue, ",")

		// fmt.Println(next)

		nextI, err := strconv.Atoi(next[0])
		nextJ, err := strconv.Atoi(next[1])

		if err != nil {
			fmt.Println(err)
		}

		i = nextI
		j = nextJ

		queue = queue[1:]
	}

	fmt.Println(C)
}
