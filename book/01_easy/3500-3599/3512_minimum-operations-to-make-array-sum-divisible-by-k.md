# 3512 — Minimum Operations To Make Array Sum Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumOperationsToMakeArraySumDivisibleByK(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3512: Minimum Operations to Make Array Sum Divisible by K
// https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeArraySumDivisibleByK([]int{3, 9, 7}, 5))
	fmt.Println(MinimumOperationsToMakeArraySumDivisibleByK([]int{4, 1, 3}, 4))
}

// MinimumOperationsToMakeArraySumDivisibleByK returns the min operations (incrementing elements by 1) to make sum divisible by k.
// Time: O(n). Space: O(1).
func MinimumOperationsToMakeArraySumDivisibleByK(nums []int, k int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	rem := sum % k
	if rem == 0 {
		return 0
	}
	return rem
}
```
