# 3178 — Find The Child Who Has The Ball After K Seconds

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheChildWhoHasTheBallAfterKSeconds(n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3178: Find the Child Who Has the Ball After K Seconds
// https://leetcode.com/problems/find-the-child-who-has-the-ball-after-k-seconds/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheChildWhoHasTheBallAfterKSeconds(3, 5))
	fmt.Println(FindTheChildWhoHasTheBallAfterKSeconds(5, 6))
	fmt.Println(FindTheChildWhoHasTheBallAfterKSeconds(4, 2))
}

// FindTheChildWhoHasTheBallAfterKSeconds returns the child who has the ball after k seconds.
// Children pass the ball left-to-right, then right-to-left, repeatedly.
// Time: O(1). Space: O(1).
func FindTheChildWhoHasTheBallAfterKSeconds(n int, k int) int {
	cycleLen := 2 * (n - 1)
	k %= cycleLen
	if k < n {
		return k
	}
	return cycleLen - k
}
```
