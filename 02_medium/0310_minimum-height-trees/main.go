package main

// LeetCode #310: Minimum Height Trees
// https://leetcode.com/problems/minimum-height-trees/
// Difficulty: Medium
// Time: O(V), Space: O(V+E)

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adj := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}

	for _, edge := range edges {
		adj[edge[0]][edge[1]] = true
		adj[edge[1]][edge[0]] = true
	}

	leaves := []int{}
	for i := 0; i < n; i++ {
		if len(adj[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	remaining := n
	for remaining > 2 {
		remaining -= len(leaves)
		newLeaves := []int{}

		for _, leaf := range leaves {
			for neighbor := range adj[leaf] {
				delete(adj[neighbor], leaf)
				if len(adj[neighbor]) == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}

		leaves = newLeaves
	}

	return leaves
}

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	fmt.Println(findMinHeightTrees(1, [][]int{}))
}
