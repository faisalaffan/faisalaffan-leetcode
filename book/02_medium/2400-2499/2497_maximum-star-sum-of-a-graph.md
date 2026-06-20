# 2497 — Maximum Star Sum Of A Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func maxStarSum(vals []int, edges [][]int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n + m log k)  |  **Ruang:** O(n + m)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2497: Maximum Star Sum of a Graph
// https://leetcode.com/problems/maximum-star-sum-of-a-graph/
// Difficulty: Medium
// Time: O(n + m log k) | Space: O(n + m)
// For each node, sort neighbor values descending, take top k positive.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxStarSum([]int{1, 2, 3, 4, 10, -10, -20}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}, {3, 6}}, 2))
	// 16

	fmt.Println(maxStarSum([]int{-5}, [][]int{}, 0))
	// -5
}

func maxStarSum(vals []int, edges [][]int, k int) int {
	n := len(vals)
  // Matriks 2D
	neighbors := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		neighbors[u] = append(neighbors[u], vals[v])
		neighbors[v] = append(neighbors[v], vals[u])
	}

	ans := vals[0]
	for i := 0; i < n; i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(neighbors[i])))
		sum := vals[i]
		for j := 0; j < k && j < len(neighbors[i]); j++ {
			if neighbors[i][j] > 0 {
				sum += neighbors[i][j]
			}
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
```
