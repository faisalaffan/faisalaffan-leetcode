# 3175 — Find The First Player To Win K Games In A Row

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findWinningPlayer(skills []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3175: Find The First Player to Win K Games in a Row
// https://leetcode.com/problems/find-the-first-player-to-win-k-games-in-a-row/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func findWinningPlayer(skills []int, k int) int {
	n := len(skills)
	maxIdx := 0
	curWins := 0

	for i := 1; i < n; i++ {
		if skills[i] > skills[maxIdx] {
			maxIdx = i
			curWins = 1
		} else {
			curWins++
		}
		if curWins >= k {
			return maxIdx
		}
	}
	return maxIdx
}

func main() {
	fmt.Println(findWinningPlayer([]int{4, 2, 6, 3, 9}, 2)) // Expected: 2
	fmt.Println(findWinningPlayer([]int{2, 5, 4}, 3))        // Expected: 1
}
```
