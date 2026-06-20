# 2374 — Node With Highest Edge Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func edgeScore(edges []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2374: Node With Highest Edge Score
// https://leetcode.com/problems/node-with-highest-edge-score/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Each node i points to edges[i]. Score of j = sum of i where edges[i] == j.

import "fmt"

func main() {
	fmt.Println(edgeScore([]int{1, 0, 0, 0, 0, 7, 7, 5})) // 7
	fmt.Println(edgeScore([]int{2, 0, 0, 2}))               // 0
}

func edgeScore(edges []int) int {
	n := len(edges)
  // Alokasi slice integer
	score := make([]int, n)
	for i, to := range edges {
		score[to] += i
	}

	maxScore := -1
	ans := -1
	for i, s := range score {
		if s > maxScore {
			maxScore = s
			ans = i
		}
	}
	return ans
}
```
