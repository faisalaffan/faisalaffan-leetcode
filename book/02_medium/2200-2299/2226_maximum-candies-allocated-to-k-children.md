# 2226 — Maximum Candies Allocated To K Children

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumCandies(candies []int, k int64) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2226: Maximum Candies Allocated to K Children
// https://leetcode.com/problems/maximum-candies-allocated-to-k-children/
// Difficulty: Medium
// Time: O(n log m) | Space: O(1)

import "fmt"

func maximumCandies(candies []int, k int64) int {
	lo, hi := 1, 0
	for _, c := range candies {
		if c > hi {
			hi = c
		}
	}

	result := 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		var count int64 = 0
		for _, c := range candies {
			count += int64(c / mid)
		}
		if count >= k {
			result = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(maximumCandies([]int{5, 8, 6}, 3))
	// Expected: 5

	// Test case 2
	fmt.Println(maximumCandies([]int{2, 5}, 11))
	// Expected: 0

	// Test case 3
	fmt.Println(maximumCandies([]int{4, 7, 5}, 16))
	// Expected: 1
}
```
