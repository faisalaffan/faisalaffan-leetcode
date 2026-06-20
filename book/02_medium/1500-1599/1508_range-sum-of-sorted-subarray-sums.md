# 1508 — Range Sum Of Sorted Subarray Sums

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RangeSum(nums []int, n int, left int, right int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(N^2 log N), Space: O(N^2)  
**Kompleksitas Ruang:** O(N^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1508: Range Sum of Sorted Subarray Sums
// https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
	fmt.Println(RangeSum([]int{1, 2, 3, 4}, 4, 3, 4))
	fmt.Println(RangeSum([]int{1, 2, 3, 4}, 4, 1, 10))
}

func RangeSum(nums []int, n int, left int, right int) int {
	// Time: O(N^2 log N), Space: O(N^2)
	const mod = 1_000_000_007

	// Generate all subarray sums
  // Alokasi slice integer
	sums := make([]int, 0, n*(n+1)/2)
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			sums = append(sums, sum)
		}
	}

	// Sort
  // Urutkan secara ascending — O(n log n)
	sort.Ints(sums)

	// Sum from left-1 to right-1
	result := 0
	for i := left - 1; i < right; i++ {
		result = (result + sums[i]) % mod
	}

	return result
}
```
