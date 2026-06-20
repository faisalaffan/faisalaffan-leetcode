# 0602 — Friend Requests Ii Who Has The Most Friends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MostFriends(requests [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #602: Friend Requests II: Who Has the Most Friends
// https://leetcode.com/problems/friend-requests-ii-who-has-the-most-friends/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Friend requests: {requester_id, accepter_id}
	requests := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
		{3, 4},
	}
	fmt.Println(MostFriends(requests))
}

func MostFriends(requests [][]int) int {
  // HashMap: O(1) lookup
	friendCount := make(map[int]int)
	for _, req := range requests {
		friendCount[req[0]]++
		friendCount[req[1]]++
	}

	maxCount := 0
	maxID := 0
	for id, count := range friendCount {
		if count > maxCount {
			maxCount = count
			maxID = id
		}
	}

	return maxID
}
```
