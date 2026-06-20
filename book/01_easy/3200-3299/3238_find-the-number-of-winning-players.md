# 3238 — Find The Number Of Winning Players

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheNumberOfWinningPlayers(n int, pick [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3238: Find the Number of Winning Players
// https://leetcode.com/problems/find-the-number-of-winning-players/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheNumberOfWinningPlayers(4, [][]int{{0, 0}, {1, 0}, {1, 0}, {2, 1}, {2, 1}, {2, 0}}))
	fmt.Println(FindTheNumberOfWinningPlayers(5, [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}}))
}

// FindTheNumberOfWinningPlayers counts players who have picked at least i+1 balls of the same color (where i is player index).
// Time: O(n). Space: O(n).
func FindTheNumberOfWinningPlayers(n int, pick [][]int) int {
	// Count colors per player
  // Alokasi slice integer
	playerColors := make([]map[int]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range playerColors {
		playerColors[i] = make(map[int]int)
	}
	for _, p := range pick {
		player, color := p[0], p[1]
		playerColors[player][color]++
	}

	winners := 0
	for i := 0; i < n; i++ {
		for _, count := range playerColors[i] {
			if count > i {
				winners++
				break
			}
		}
	}
	return winners
}
```
