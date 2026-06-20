# 1823 — Find The Winner Of The Circular Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findTheWinner(n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1823: Find the Winner of the Circular Game
// https://leetcode.com/problems/find-the-winner-of-the-circular-game/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func findTheWinner(n int, k int) int {
	winner := 0 // 0-indexed position for 1 person
	for i := 2; i <= n; i++ {
		winner = (winner + k) % i
	}
	return winner + 1 // convert to 1-indexed
}

func main() {
	fmt.Println(findTheWinner(5, 2)) // Expected: 3
	fmt.Println(findTheWinner(6, 5)) // Expected: 1
	fmt.Println(findTheWinner(1, 1)) // Expected: 1
}
```
