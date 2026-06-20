# 3824 — Minimum K To Reduce Array Within Limit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumKToReduceArrayWithinLimit(nums []int) int
```

> **💡 Hint:** Binary search on k. Check if total operations <= k^2.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(N log M) where M = max(nums)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3824: Minimum K to Reduce Array Within Limit
// https://leetcode.com/problems/minimum-k-to-reduce-array-within-limit/
// Difficulty: Medium
// Time: O(N log M) where M = max(nums) | Space: O(1)
// Approach: Binary search on k. Check if total operations <= k^2.

import "fmt"

func MinimumKToReduceArrayWithinLimit(nums []int) int {
	// Binary search on k
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	lo, hi := 1, maxVal
	ans := maxVal

	for lo <= hi {
		mid := (lo + hi) / 2
		ops := 0
		for _, v := range nums {
			ops += (v + mid - 1) / mid // ceil(v / mid)
		}
		if ops <= mid*mid {
			ans = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumKToReduceArrayWithinLimit([]int{3, 7, 5})) // Expected: 3

	// Example 2
	fmt.Println(MinimumKToReduceArrayWithinLimit([]int{1})) // Expected: 1

	// Example 3
	fmt.Println(MinimumKToReduceArrayWithinLimit([]int{10, 10, 10}))
}
```
