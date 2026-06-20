# 2225 — Find Players With Zero Or One Losses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findWinners(matches [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2225: Find Players With Zero or One Losses
// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
  // Membuat map (HashMap) — pencarian O(1)
	losses := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	players := make(map[int]bool)

	for _, m := range matches {
		winner, loser := m[0], m[1]
		players[winner] = true
		players[loser] = true
		losses[loser]++
	}

	winners := []int{}
	oneLoss := []int{}
	for p := range players {
		l := losses[p]
		if l == 0 {
			winners = append(winners, p)
		} else if l == 1 {
			oneLoss = append(oneLoss, p)
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(winners)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(oneLoss)

	return [][]int{winners, oneLoss}
}

func main() {
	// Test case 1
	fmt.Println(findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	// Expected: [[1,2,10],[4,5,7,8]]

	// Test case 2
	fmt.Println(findWinners([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}))
	// Expected: [[1,2,5,6],[]]
}
```
