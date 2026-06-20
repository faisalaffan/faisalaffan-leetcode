# 3558 — Number Of Ways To Assign Edge Weights I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func NumberOfWaysToAssignEdgeWeightsI(n int, edges [][]int, weights []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3558: Number of Ways to Assign Edge Weights I
// https://leetcode.com/problems/number-of-ways-to-assign-edge-weights-i/
// Difficulty: Medium
// Complexity: O(n * w) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	n := 3
	edges := [][]int{{0, 1}, {1, 2}}
	weights := []int{1, 2}
	fmt.Println("Test 1:", NumberOfWaysToAssignEdgeWeightsI(n, edges, weights))
	// Test case 2
	n2 := 4
	edges2 := [][]int{{0, 1}, {0, 2}, {0, 3}}
	weights2 := []int{1, 1, 2}
	fmt.Println("Test 2:", NumberOfWaysToAssignEdgeWeightsI(n2, edges2, weights2))
	// Test case 3
	n3 := 2
	edges3 := [][]int{{0, 1}}
	weights3 := []int{5}
	fmt.Println("Test 3:", NumberOfWaysToAssignEdgeWeightsI(n3, edges3, weights3))
}

func NumberOfWaysToAssignEdgeWeightsI(n int, edges [][]int, weights []int) int {
	mod := 1000000007
	// Count how many ways to assign each weight to an edge
	// Simple case: each weight can go to any edge
	if len(edges) == 0 || len(weights) == 0 {
		return 1
	}
	if len(edges) != len(weights) {
		return 0
	}
	// Number of permutations of weights assigned to edges
	result := 1
	for i := 2; i <= len(weights); i++ {
		result = (result * i) % mod
	}
	return result
}
```
