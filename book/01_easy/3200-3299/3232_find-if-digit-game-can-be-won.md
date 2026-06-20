# 3232 — Find If Digit Game Can Be Won

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindIfDigitGameCanBeWon(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3232: Find if Digit Game Can Be Won
// https://leetcode.com/problems/find-if-digit-game-can-be-won/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindIfDigitGameCanBeWon([]int{1, 2, 3, 4, 10}))
	fmt.Println(FindIfDigitGameCanBeWon([]int{1, 2, 3, 4, 5, 14}))
}

// FindIfDigitGameCanBeWon returns true if Alice can win the digit game.
// Alice takes all single-digit numbers, Bob takes all two-digit+ numbers.
// Alice wins if the sums are not equal.
// Time: O(n). Space: O(1).
func FindIfDigitGameCanBeWon(nums []int) bool {
	sumSingle := 0
	sumDouble := 0
	for _, num := range nums {
		if num < 10 {
			sumSingle += num
		} else {
			sumDouble += num
		}
	}
	return sumSingle != sumDouble
}
```
