# 3663 — Find The Least Frequent Digit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheLeastFrequentDigit(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n) - number of digits  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3663: Find The Least Frequent Digit
// https://leetcode.com/problems/find-the-least-frequent-digit/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindTheLeastFrequentDigit(1553322))
	fmt.Println(FindTheLeastFrequentDigit(723344511))
}

// Time: O(log n) - number of digits
// Space: O(1)
func FindTheLeastFrequentDigit(n int) int {
	cnt := [10]int{}
	for n > 0 {
		cnt[n%10]++
		n /= 10
	}

	minCnt := math.MaxInt
	ans := 0
	for d := 0; d <= 9; d++ {
		if cnt[d] > 0 && (cnt[d] < minCnt || (cnt[d] == minCnt && d < ans)) {
			minCnt = cnt[d]
			ans = d
		}
	}
	return ans
}
```
