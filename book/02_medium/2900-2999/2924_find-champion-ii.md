# 2924 — Find Champion Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func findChampionII(n int, edges [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n+m)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2924: Find Champion II
// https://leetcode.com/problems/find-champion-ii/
// Difficulty: Medium
// Time: O(n+m) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(findChampionII(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findChampionII(4, [][]int{{0, 2}, {1, 3}, {1, 2}}))
	fmt.Println(findChampionII(2, [][]int{{0, 1}}))
}

func findChampionII(n int, edges [][]int) int {
  // Alokasi slice
	indeg := make([]int, n)
	for _, e := range edges {
		indeg[e[1]]++
	}
	ans, cnt := -1, 0
	for i, x := range indeg {
		if x == 0 {
			cnt++
			ans = i
		}
	}
	if cnt == 1 {
		return ans
	}
	return -1
}
```
