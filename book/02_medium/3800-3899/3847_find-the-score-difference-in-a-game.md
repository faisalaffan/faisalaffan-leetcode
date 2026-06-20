# 3847 — Find The Score Difference In A Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheScoreDifferenceInAGame(nums []int) int
```

> **💡 Hint:** Track active player and swap on odd or every 6th game.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3847: Find the Score Difference in a Game
// https://leetcode.com/problems/find-the-score-difference-in-a-game/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Track active player and swap on odd or every 6th game.

import "fmt"

func FindTheScoreDifferenceInAGame(nums []int) int {
	first, second := 0, 0
	activeIsFirst := true

	for i, v := range nums {
		// Swap if odd
		if v%2 == 1 {
			activeIsFirst = !activeIsFirst
		}
		// Swap every 6th game (0-indexed, so i%6 == 5)
		if i%6 == 5 {
			activeIsFirst = !activeIsFirst
		}
		if activeIsFirst {
			first += v
		} else {
			second += v
		}
	}

	return first - second
}

func main() {
	// Example 1
	fmt.Println(FindTheScoreDifferenceInAGame([]int{1, 2, 3})) // Expected: 0

	// Example 2
	fmt.Println(FindTheScoreDifferenceInAGame([]int{2, 4, 2, 1, 2, 1})) // Expected: 4

	// Example 3
	fmt.Println(FindTheScoreDifferenceInAGame([]int{1})) // Expected: -1
}
```
