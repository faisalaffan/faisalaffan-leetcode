# 0090 — Subsets Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func subsetsWithDup(nums []int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * 2^n)  
**Kompleksitas Ruang:** O(n * 2^n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #90: Subsets II
// https://leetcode.com/problems/subsets-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	result := [][]int{{}}
	start := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		n := len(result)
		begin := 0
		if i > 0 && nums[i] == nums[i-1] {
			begin = start
		}
		start = n
		for j := begin; j < n; j++ {
  // Alokasi slice integer
			newSubset := make([]int, len(result[j])+1)
			copy(newSubset, result[j])
			newSubset[len(result[j])] = nums[i]
			result = append(result, newSubset)
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	// [[] [1] [2] [1 2] [2 2] [1 2 2]]

	// Test case 2
	fmt.Println(subsetsWithDup([]int{0})) // [[] [0]]
}

// Time: O(n * 2^n) | Space: O(n * 2^n)
```
