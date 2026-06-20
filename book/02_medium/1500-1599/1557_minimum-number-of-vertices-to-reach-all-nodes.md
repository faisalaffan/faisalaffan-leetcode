# 1557 — Minimum Number Of Vertices To Reach All Nodes

## Deskripsi

**Soal:** [1557. Minimum Number Of Vertices To Reach All Nodes](https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N + E), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1557: Minimum Number of Vertices to Reach All Nodes
// https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindSmallestSetOfVertices(6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}))
	fmt.Println(FindSmallestSetOfVertices(3, [][]int{{0, 1}, {2, 1}}))
	fmt.Println(FindSmallestSetOfVertices(5, [][]int{{0, 1}, {2, 1}, {3, 1}, {4, 0}}))
}

func FindSmallestSetOfVertices(n int, edges [][]int) []int {
	// Time: O(N + E), Space: O(N)
	// Nodes with indegree 0 must be in the result since they can't be reached
  // Membuat slice untuk menyimpan hasil
	indegree := make([]int, n)
	for _, e := range edges {
		indegree[e[1]]++
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			result = append(result, i)
		}
	}

	return result
}
```
