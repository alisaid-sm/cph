package main

import (
	"fmt"
)

/*
	1: 2,3
	2: 1
	3: 1

	1: 2
	2: 1
	3: 0

	1: true
	2: false
	3: false

    O(N*M)
*/

func main() {
	var N, M int

	fmt.Scan(&N, &M)

	playerTargetMap := make(map[int][]int)
	match := make([]int, N)
	visited := make([]bool, N)

	for i := 0; i < M; i++ {
		var player int
		var target int
		fmt.Scan(&player, &target)

		playerTargetMap[player] = append(playerTargetMap[player], target)
		playerTargetMap[target] = append(playerTargetMap[target], player)
	}

	var dfs func(int) bool
	dfs = func(p int) bool {
		for _, target := range playerTargetMap[p] {
			if visited[target-1] {
				continue
			}
			visited[target-1] = true

			if match[target-1] == 0 || dfs(match[target-1]) {
				match[target-1] = p

				return true
			}
		}

		return false
	}

	var matchCount int

	for i := 0; i < N; i++ {
		if dfs(i + 1) {

			matchCount++
		}
		for j := range visited {
			visited[j] = false
		}
	}

	if matchCount != N {
		fmt.Println("Impossible")
		return
	}

	for _, matched := range match {
		fmt.Println(matched)
	}
}
