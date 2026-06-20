# 3759 — Count Elements With At Least K Greater Values

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countElementsWithAtLeastKGreaterValues(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3759: Count Elements With at Least K Greater Values
// https://leetcode.com/problems/count-elements-with-at-least-k-greater-values/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func countElementsWithAtLeastKGreaterValues(nums []int, k int) int {
	if k == 0 {
		return len(nums)
	}
	n := len(nums)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	left := n - k
	// Skip duplicates of threshold value
	for left-1 >= 0 && nums[left-1] == nums[left] {
		left--
	}
	return left
}

func main() {
	fmt.Println(countElementsWithAtLeastKGreaterValues([]int{3, 1, 2}, 1))
	fmt.Println(countElementsWithAtLeastKGreaterValues([]int{5, 5, 5}, 2))
	fmt.Println(countElementsWithAtLeastKGreaterValues([]int{1, 2, 3, 4, 5}, 2))
}
```
