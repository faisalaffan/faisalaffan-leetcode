# 2483 — Minimum Penalty For A Shop

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func bestClosingTime(customers string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2483: Minimum Penalty for a Shop
// https://leetcode.com/problems/minimum-penalty-for-a-shop/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Penalty at hour i = 'N' before i + 'Y' after/at i. Sweep left to right.

import "fmt"

func main() {
	fmt.Println(bestClosingTime("YYNY")) // 2
	fmt.Println(bestClosingTime("NNNNN")) // 0
	fmt.Println(bestClosingTime("YYYY"))  // 4
}

func bestClosingTime(customers string) int {
	// Start: close at hour 0
	penalty := 0
	for _, ch := range customers {
		if ch == 'Y' {
			penalty++
		}
	}
	minPenalty := penalty
	bestHour := 0

	// Try closing at hour 1..n
	for i, ch := range customers {
		if ch == 'Y' {
			penalty--
		} else {
			penalty++
		}
		if penalty < minPenalty {
			minPenalty = penalty
			bestHour = i + 1
		}
	}
	return bestHour
}
```
