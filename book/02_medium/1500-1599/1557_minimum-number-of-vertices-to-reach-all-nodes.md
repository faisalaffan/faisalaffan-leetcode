# 1557 — Minimum Number Of Vertices To Reach All Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func FindSmallestSetOfVertices(n int, edges [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N + E), Space: O(N)  |  **Ruang:** O(N)


## 💻 Solusi Go

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
  // Alokasi slice
	indegree := make([]int, n)
	for _, e := range edges {
		indegree[e[1]]++
	}

  // Alokasi slice
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			result = append(result, i)
		}
	}

	return result
}
```
