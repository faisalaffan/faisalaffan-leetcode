# 0614 — Second Degree Follower

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func SecondDegreeFollowers(follows [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #614: Second Degree Follower
// https://leetcode.com/problems/second-degree-follower/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Follow relationships: {follower, followee}
	follows := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 1},
		{3, 2},
		{4, 5},
	}
	fmt.Println(SecondDegreeFollowers(follows))
}

func SecondDegreeFollowers(follows [][]int) []int {
  // HashMap: O(1) lookup
	followers := make(map[int]map[int]bool)

	for _, f := range follows {
		follower, followee := f[0], f[1]
		if followers[followee] == nil {
			followers[followee] = make(map[int]bool)
		}
		followers[followee][follower] = true
	}

	result := []int{}
	for userID, fMap := range followers {
		if len(fMap) >= 2 {
			result = append(result, userID)
		}
	}

	return result
}
```
