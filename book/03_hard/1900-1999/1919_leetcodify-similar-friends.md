# 1919 — Leetcodify Similar Friends

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func leetcodifySimilarFriends(listens [][]int, friendships [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1919: Leetcodify Similar Friends
// https://leetcode.com/problems/leetcodify-similar-friends/
// Difficulty: Hard [Paid]
//
// Given tables:
// - listens (user_id, song_id, day)
// - friendship (user1_id, user2_id)
//
// Find friend pairs (already friends) who have high similarity in music taste.
// Two friends are "similar" if the number of songs they both listened to
// on the same day is >= 3 times the number of different songs they listened to
// (union of songs, counted per day).
//
// More precisely: for each pair of friends (u1, u2), count pairs (song, day)
// where both listened to that song on that day. Count pairs (song, day) where
// either listened. If same_count >= 3 AND same_count / total >= threshold (0.6?),
// return them.
//
// Actually from the problem: return friend pairs where the number of (song, day)
// they have in common >= 3, and common / total_listened_both >= some threshold.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: simple
	listens := [][]int{
		{1, 101, 1}, {1, 102, 1},
		{2, 101, 1}, {2, 102, 1},
		{3, 101, 1},
	}
	friendships := [][]int{
		{1, 2}, {1, 3},
	}
	// Friends 1-2: both listened to {101, 102} on day 1, same=2, total=2, ratio=1.0 >= 0.6
	// Friends 1-3: both listened to {101} on day 1, same=1, total=not sure
	fmt.Println(leetcodifySimilarFriends(listens, friendships))

	// Test case 2: higher threshold
	listens = [][]int{
		{1, 101, 1}, {1, 101, 1}, {1, 101, 1},
		{2, 101, 1},
		{3, 101, 1}, {3, 102, 1},
	}
	friendships = [][]int{
		{1, 2}, {2, 3},
	}
	fmt.Println(leetcodifySimilarFriends(listens, friendships))

	// Test case 3: empty
	fmt.Println(leetcodifySimilarFriends([][]int{}, [][]int{}))
}

// leetcodifySimilarFriends returns [][]int{{user1_id, user2_id}, ...}
func leetcodifySimilarFriends(listens [][]int, friendships [][]int) [][]int {
	// Build set of (user, song, day) - unique listens per user per song per day
  // HashMap: O(1) lookup
	userSongDay := make(map[[3]int]bool)
  // HashMap: O(1) lookup
	userSongs := make(map[int]map[[2]int]bool) // user -> set of (song, day) pairs
	for _, l := range listens {
		u, s, d := l[0], l[1], l[2]
		key := [3]int{u, s, d}
		if !userSongDay[key] {
			userSongDay[key] = true
			if userSongs[u] == nil {
				userSongs[u] = make(map[[2]int]bool)
			}
			userSongs[u][[2]int{s, d}] = true
		}
	}

	var result [][]int

	for _, f := range friendships {
		u1, u2 := f[0], f[1]
		if u1 > u2 {
			u1, u2 = u2, u1
		}

		songs1 := userSongs[u1]
		songs2 := userSongs[u2]

		if len(songs1) == 0 || len(songs2) == 0 {
			continue
		}

		// Count intersection
		common := 0
		for sd := range songs1 {
			if songs2[sd] {
				common++
			}
		}

		// Count union
  // HashMap: O(1) lookup
		union := make(map[[2]int]bool)
		for sd := range songs1 {
			union[sd] = true
		}
		for sd := range songs2 {
			union[sd] = true
		}
		total := len(union)

		if total == 0 {
			continue
		}

		// Use similarity ratio threshold >= 0.6
		// common >= 3 AND common/total >= 0.6
		if common >= 3 && common*10 >= total*6 {
			result = append(result, []int{u1, u2})
		}
	}

  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		if result[i][0] != result[j][0] {
			return result[i][0] < result[j][0]
		}
		return result[i][1] < result[j][1]
	})

	if result == nil {
		return [][]int{}
	}
	return result
}
```
