# 3560 — Find Minimum Log Transportation Cost

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindMinimumLogTransportationCost(n int, m int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3560: Find Minimum Log Transportation Cost
// https://leetcode.com/problems/find-minimum-log-transportation-cost/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMinimumLogTransportationCost(10, 10, 3))
	fmt.Println(FindMinimumLogTransportationCost(5, 5, 2))
}

// FindMinimumLogTransportationCost returns the minimum transportation cost for logs of size n x m, cutting with factor k.
// Cost = max(0, max(n, m) - k) * k
// Time: O(1). Space: O(1).
func FindMinimumLogTransportationCost(n int, m int, k int) int {
	larger := n
	if m > larger {
		larger = m
	}
	if larger <= k {
		return 0
	}
	return (larger - k) * k
}
```
