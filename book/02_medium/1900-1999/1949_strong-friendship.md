# 1949 — Strong Friendship

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func StrongFriendship(friendships [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^2), Space: O(n^2)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1949: Strong Friendship
// https://leetcode.com/problems/strong-friendship/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// friendships: [user1_id, user2_id]
	friendships := [][]int{{1, 2}, {1, 3}, {2, 3}, {1, 4}, {2, 4}, {1, 5}}
	fmt.Println(StrongFriendship(friendships))
}

// Time: O(n^2), Space: O(n^2)
func StrongFriendship(friendships [][]int) int {
  // HashMap: O(1) lookup
	friendSet := make(map[int]map[int]bool)
	for _, f := range friendships {
		a, b := f[0], f[1]
		if friendSet[a] == nil {
			friendSet[a] = make(map[int]bool)
		}
		if friendSet[b] == nil {
			friendSet[b] = make(map[int]bool)
		}
		friendSet[a][b] = true
		friendSet[b][a] = true
	}

	count := 0
	// For each pair of users, check if they have at least 3 common friends
  // Alokasi slice
	users := make([]int, 0, len(friendSet))
	for u := range friendSet {
		users = append(users, u)
	}

  // Linear scan O(n)
	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			a, b := users[i], users[j]
			if friendSet[a][b] { // they are friends
				continue
			}
			common := 0
			for f := range friendSet[a] {
				if friendSet[b][f] {
					common++
				}
			}
			if common >= 3 {
				count++
			}
		}
	}
	return count
}
```
