# 2750 — Ways To Split Array Into Good Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func WaysToSplitArrayIntoGoodSubarrays(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2750: Ways to Split Array Into Good Subarrays
// https://leetcode.com/problems/ways-to-split-array-into-good-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func WaysToSplitArrayIntoGoodSubarrays(nums []int) int {
	// Find positions of 1s
  // Alokasi slice integer
	ones := make([]int, 0)
	for i, v := range nums {
		if v == 1 {
			ones = append(ones, i)
		}
	}

	if len(ones) == 0 {
		return 0
	}

	const mod = 1_000_000_007
	result := 1
	for i := 1; i < len(ones); i++ {
		gap := ones[i] - ones[i-1]
		result = (result * gap) % mod
	}

	return result
}

func main() {
	fmt.Println(WaysToSplitArrayIntoGoodSubarrays([]int{0, 1, 0, 0, 1}))
	fmt.Println(WaysToSplitArrayIntoGoodSubarrays([]int{1, 1, 1}))
	fmt.Println(WaysToSplitArrayIntoGoodSubarrays([]int{0, 0, 0}))
}
```
