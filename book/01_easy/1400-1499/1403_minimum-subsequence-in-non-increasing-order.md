# 1403 — Minimum Subsequence In Non Increasing Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minSubsequence(nums []int) []int

import (
	"fmt"
	"sort"
)

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1) excluding output  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1403: Minimum Subsequence in Non-Increasing Order
// https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order/
// Difficulty: Easy
//
// LeetCode submission: func minSubsequence(nums []int) []int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumSubsequenceInNonIncreasingOrder([]int{4, 3, 10, 9, 8})) // [10 9]
	fmt.Println(MinimumSubsequenceInNonIncreasingOrder([]int{4, 4, 7, 6, 7}))  // [7 7 6]
}

// Time: O(n log n), Space: O(1) excluding output
func MinimumSubsequenceInNonIncreasingOrder(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		sum += v
		if sum > total-sum {
			return nums[:i+1]
		}
	}
	return nums
}
```
