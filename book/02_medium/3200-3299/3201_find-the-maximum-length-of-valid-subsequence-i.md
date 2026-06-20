# 3201 — Find The Maximum Length Of Valid Subsequence I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumLength(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3201: Find the Maximum Length of Valid Subsequence I
// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-i/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumLength(nums []int) int {
	// Count numbers by parity
	odd, even := 0, 0
	for _, v := range nums {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	// All odd or all even
	ans := max(odd, even)

	// Alternating parity (two patterns: even-odd-even-... or odd-even-odd-...)
	// Both give the same count: min that alternates, but since we need
	// pattern e,o,e,o,... or o,e,o,e,... starting from both parity,
	// the max alternating length is the max of:
	// For each possible (a%2, b%2) pattern where (a+b)%2 == 1 (different parity)
	// Actually the requirement is that for adjacent pairs, (a+b)%2 == 1
	// So we can alternate: 0,1,0,1,... or 1,0,1,0,...

	// Count alternating starting with even
	altEven := 0
	last := 1 // start with expecting even (0)
	for _, v := range nums {
		if v%2 == 0 && last != 0 {
			altEven++
			last = 0
		} else if v%2 == 1 && last != 1 {
			altEven++
			last = 1
		}
	}

	// Count alternating starting with odd
	altOdd := 0
	last = 0 // start with expecting odd (1)
	for _, v := range nums {
		if v%2 == 1 && last != 1 {
			altOdd++
			last = 1
		} else if v%2 == 0 && last != 0 {
			altOdd++
			last = 0
		}
	}

	ans = max(ans, max(altEven, altOdd))
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 3, 4})) // Expected: 4
	fmt.Println(maximumLength([]int{1, 3, 5}))    // Expected: 3
	fmt.Println(maximumLength([]int{2, 4, 6}))    // Expected: 3
}
```
