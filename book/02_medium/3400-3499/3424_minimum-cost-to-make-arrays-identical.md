# 3424 — Minimum Cost To Make Arrays Identical

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minCost(arr []int, brr []int, k int64) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3424: Minimum Cost to Make Arrays Identical
// https://leetcode.com/problems/minimum-cost-to-make-arrays-identical/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func minCost(arr []int, brr []int, k int64) int64 {
	var cost1 int64
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(arr); i++ {
		diff := arr[i] - brr[i]
		if diff < 0 {
			diff = -diff
		}
		cost1 += int64(diff)
	}

  // Alokasi slice integer
	sortedArr := make([]int, len(arr))
  // Alokasi slice integer
	sortedBrr := make([]int, len(brr))
	copy(sortedArr, arr)
	copy(sortedBrr, brr)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sortedArr)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sortedBrr)

	var cost2 int64
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(sortedArr); i++ {
		diff := sortedArr[i] - sortedBrr[i]
		if diff < 0 {
			diff = -diff
		}
		cost2 += int64(diff)
	}
	cost2 += k

	if cost1 < cost2 {
		return cost1
	}
	return cost2
}

func main() {
	fmt.Println(minCost([]int{4, 2, 5}, []int{6, 3, 1}, 2)) // 7
	fmt.Println(minCost([]int{1, 2, 3}, []int{4, 5, 6}, 1)) // 9
}
```
