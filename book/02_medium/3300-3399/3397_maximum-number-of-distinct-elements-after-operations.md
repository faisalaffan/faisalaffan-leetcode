# 3397 — Maximum Number Of Distinct Elements After Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxDistinctElements(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3397: Maximum Number of Distinct Elements After Operations
// https://leetcode.com/problems/maximum-number-of-distinct-elements-after-operations/
// Difficulty: Medium
// Time: O(n log n) Space: O(1)

import (
	"fmt"
	"math"
	"slices"
)

func maxDistinctElements(nums []int, k int) int {
	n := len(nums)
	if k*2+1 >= n {
		return n
	}

	slices.Sort(nums)
	pre := math.MinInt
	ans := 0
	for _, x := range nums {
		candidate := max(x-k, pre+1)
		if candidate <= x+k {
			ans++
			pre = candidate
		}
	}
	return ans
}

func main() {
	fmt.Println(maxDistinctElements([]int{1, 2, 2, 3, 3, 4}, 2)) // 6
	fmt.Println(maxDistinctElements([]int{4, 4, 4, 4}, 1))        // 3
	fmt.Println(maxDistinctElements([]int{1, 1, 1, 1}, 0))        // 1
}
```
