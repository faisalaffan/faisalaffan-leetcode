# 3821 — Find Nth Smallest Integer With K One Bits

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func nthSmallest(n int64, k int) int64
```

> **💡 Hint:** Use combinatorial ranking. Generate numbers with k bits

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3821: Find Nth Smallest Integer With K One Bits
// https://leetcode.com/problems/find-nth-smallest-integer-with-k-one-bits/
// Difficulty: Hard
//
// Find the n-th smallest positive integer that has exactly k set
// bits in its binary representation.
//
// Approach: Use combinatorial ranking. Generate numbers with k bits
// in increasing order using next combination (Gosper's hack).

import "fmt"

func main() {
	// Example 1
	fmt.Println(nthSmallest(3, 2))
	// Example 2
	fmt.Println(nthSmallest(5, 3))
	// Edge: n=1
	fmt.Println(nthSmallest(1, 1))
	// Edge: k=0
	fmt.Println(nthSmallest(1, 0))
}

func nthSmallest(n int64, k int) int64 {
	if k == 0 {
		if n == 1 {
			return 0
		}
		return -1
	}

	num := int64((1 << uint(k)) - 1)
	for i := int64(1); i < n; i++ {
		// Next number with k bits set (Gosper's hack)
		c := num & -num
		r := num + c
		num = (((r ^ num) >> 2) / c) | r
	}
	return num
}
```
