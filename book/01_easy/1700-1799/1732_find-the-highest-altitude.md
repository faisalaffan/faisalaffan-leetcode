# 1732 — Find The Highest Altitude

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func LargestAltitude(gain []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1732: Find the Highest Altitude
// https://leetcode.com/problems/find-the-highest-altitude/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func LargestAltitude(gain []int) int {
	maxAlt := 0
	current := 0
	for _, g := range gain {
		current += g
		if current > maxAlt {
			maxAlt = current
		}
	}
	return maxAlt
}

func main() {
	fmt.Println(LargestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(LargestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
```
