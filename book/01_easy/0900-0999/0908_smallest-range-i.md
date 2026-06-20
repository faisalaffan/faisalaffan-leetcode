# 0908 — Smallest Range I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func smallestRangeI(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #908: Smallest Range I
// https://leetcode.com/problems/smallest-range-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))          // 0
	fmt.Println(smallestRangeI([]int{0, 10}, 2))      // 6
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))    // 0
}

// smallestRangeI returns the smallest possible range after modifying each element by at most k.
// Time: O(n). Space: O(1).
func smallestRangeI(nums []int, k int) int {
	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	diff := (maxVal - k) - (minVal + k)
	if diff < 0 {
		return 0
	}
	return diff
}
```
