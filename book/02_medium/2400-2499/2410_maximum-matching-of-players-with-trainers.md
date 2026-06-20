# 2410 — Maximum Matching Of Players With Trainers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan aturan permainan dan harus menentukan siapa yang menang atau berapa skor maksimal. Tugasmu adalah menganalisis permainan dan membuat keputusan optimal di setiap langkah.

Soal game theory menguji kemampuanmu berpikir beberapa langkah ke depan (minimax). Seringkali diselesaikan dengan DP (Dynamic Programming) untuk menyimpan hasil subproblem.

**Konsep kunci:** minimax, optimal play, game state, DP memoization, win/lose positions.

**Fungsi yang perlu kamu implementasikan:**
```go
func matchPlayersAndTrainers(players []int, trainers []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n + m log m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2410: Maximum Matching of Players With Trainers
// https://leetcode.com/problems/maximum-matching-of-players-with-trainers/
// Difficulty: Medium
// Time: O(n log n + m log m) | Space: O(1)
// Sort both, greedy match: assign smallest sufficient trainer to each player.

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(matchPlayersAndTrainers([]int{4, 7, 9}, []int{8, 2, 5, 8})) // 2
	fmt.Println(matchPlayersAndTrainers([]int{1, 1, 1}, []int{10}))         // 1
}

func matchPlayersAndTrainers(players []int, trainers []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(players)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(trainers)
	i, j := 0, 0
	for i < len(players) && j < len(trainers) {
		if players[i] <= trainers[j] {
			i++
		}
		j++
	}
	return i
}
```
