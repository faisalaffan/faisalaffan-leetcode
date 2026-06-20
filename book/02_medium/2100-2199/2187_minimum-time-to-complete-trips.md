# 2187 — Minimum Time To Complete Trips

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumTime(time []int, totalTrips int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2187: Minimum Time to Complete Trips
// https://leetcode.com/problems/minimum-time-to-complete-trips/
// Difficulty: Medium
// Time: O(n log m) | Space: O(1)

import "fmt"

func minimumTime(time []int, totalTrips int) int64 {
	lo, hi := int64(1), int64(1)
	for _, t := range time {
		if int64(t) > hi {
			hi = int64(t)
		}
	}
	hi *= int64(totalTrips)

	for lo < hi {
		mid := lo + (hi-lo)/2
		trips := int64(0)
		for _, t := range time {
			trips += mid / int64(t)
			if trips >= int64(totalTrips) {
				break
			}
		}
		if trips >= int64(totalTrips) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	// Test case 1
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
	// Expected: 3

	// Test case 2
	fmt.Println(minimumTime([]int{2}, 1))
	// Expected: 2

	// Test case 3
	fmt.Println(minimumTime([]int{5, 10, 10}, 9))
	// Expected: 25
}
```
