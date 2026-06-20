# 2497 — Maximum Star Sum Of A Graph

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxStarSum(vals []int, edges [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m log k)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat matriks/slice 2D untuk DP
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
