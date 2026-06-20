# 2374 — Node With Highest Edge Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func edgeScore(edges []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


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
  // Alokasi slice
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
