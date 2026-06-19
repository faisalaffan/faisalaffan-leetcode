package main

// LeetCode #1345: Jump Game IV
// https://leetcode.com/problems/jump-game-iv/
// Difficulty: Hard

import "fmt"

func minJumps(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	// Build value -> indices map
	valMap := make(map[int][]int)
	for i, v := range arr {
		valMap[v] = append(valMap[v], i)
	}

	visited := make([]bool, n)
	queue := []int{0}
	visited[0] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for k := 0; k < size; k++ {
			i := queue[0]
			queue = queue[1:]

			if i == n-1 {
				return steps
			}

			// Jump i+1
			if i+1 < n && !visited[i+1] {
				visited[i+1] = true
				queue = append(queue, i+1)
			}

			// Jump i-1
			if i-1 >= 0 && !visited[i-1] {
				visited[i-1] = true
				queue = append(queue, i-1)
			}

			// Jump to same-value indices
			if indices, ok := valMap[arr[i]]; ok {
				for _, j := range indices {
					if !visited[j] {
						visited[j] = true
						queue = append(queue, j)
					}
				}
				// Clear to avoid reprocessing
				delete(valMap, arr[i])
			}
		}
		steps++
	}

	return -1
}

func main() {
	// Example 1
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
	// Expected: 3

	// Example 2
	fmt.Println(minJumps([]int{7}))
	// Expected: 0

	// Example 3
	fmt.Println(minJumps([]int{7, 6, 9, 6, 9, 6, 9, 7}))
	// Expected: 1
}
