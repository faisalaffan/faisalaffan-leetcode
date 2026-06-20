# 2049 — Count Nodes With The Highest Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func countHighestScoreNodes(parents []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2049: Count Nodes With the Highest Score
// https://leetcode.com/problems/count-nodes-with-the-highest-score/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
  // Membuat matriks/slice 2D untuk DP
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}

  // Alokasi slice integer
	subtreeSize := make([]int, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		size := 1
		for _, v := range children[u] {
			size += dfs(v)
		}
		subtreeSize[u] = size
		return size
	}
	dfs(0)

	maxScore := 0
	count := 0
	for i := 0; i < n; i++ {
		score := 1
		remaining := n
		for _, v := range children[i] {
			score *= subtreeSize[v]
			remaining -= subtreeSize[v]
		}
		if i != 0 {
			remaining = n - subtreeSize[i]
		} else {
			remaining = 0
		}
		if remaining > 0 {
			score *= remaining
		}

		if score > maxScore {
			maxScore = score
			count = 1
		} else if score == maxScore {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", countHighestScoreNodes([]int{-1, 2, 0}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", countHighestScoreNodes([]int{-1, 0, 1, 2}))
	// Expected: 1
}
```
