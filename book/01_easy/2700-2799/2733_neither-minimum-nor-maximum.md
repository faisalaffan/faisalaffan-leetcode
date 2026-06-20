# 2733 — Neither Minimum Nor Maximum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NeitherMinimumNorMaximum(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2733: Neither Minimum nor Maximum
// https://leetcode.com/problems/neither-minimum-nor-maximum/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NeitherMinimumNorMaximum([]int{3, 2, 1, 4}))
	fmt.Println(NeitherMinimumNorMaximum([]int{1, 2}))
}

func NeitherMinimumNorMaximum(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	for _, v := range nums {
		if v != minVal && v != maxVal {
			return v
		}
	}

	return -1
}
```
