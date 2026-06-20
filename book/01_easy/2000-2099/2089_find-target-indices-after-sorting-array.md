# 2089 — Find Target Indices After Sorting Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTargetIndicesAfterSortingArray(nums []int, target int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1) ignoring sort  
**Kompleksitas Ruang:** O(1) ignoring sort

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2089: Find Target Indices After Sorting Array
// https://leetcode.com/problems/find-target-indices-after-sorting-array/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 2)) // [1 2]
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 3)) // [3]
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 5)) // [4]
}

// Time: O(n log n), Space: O(1) ignoring sort
func FindTargetIndicesAfterSortingArray(nums []int, target int) []int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	var result []int
	for i, v := range nums {
		if v == target {
			result = append(result, i)
		}
	}
	return result
}
```
