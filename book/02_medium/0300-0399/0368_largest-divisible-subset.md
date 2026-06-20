# 0368 — Largest Divisible Subset

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func largestDivisibleSubset(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #368: Largest Divisible Subset
// https://leetcode.com/problems/largest-divisible-subset/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
  // Edge case: input kosong — langsung return
	if len(nums) == 0 {
		return []int{}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)
  // Alokasi slice integer
	dp := make([]int, n) // size of largest subset ending at i
  // Alokasi slice integer
	prev := make([]int, n)
	maxIdx := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIdx] {
			maxIdx = i
		}
	}

	// Reconstruct
  // Alokasi slice integer
	result := make([]int, 0, dp[maxIdx])
	for i := maxIdx; i >= 0; i = prev[i] {
		result = append(result, nums[i])
		// Reverse the traversal
	}
	// Reverse result
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", largestDivisibleSubset([]int{1, 2, 3}))
	// Expected: [1, 2] or [1, 3]

	// Test case 2
	fmt.Println("Test 2:", largestDivisibleSubset([]int{1, 2, 4, 8}))
	// Expected: [1, 2, 4, 8]

	// Test case 3
	fmt.Println("Test 3:", largestDivisibleSubset([]int{3, 4, 8, 16}))
	// Expected: [4, 8, 16] or [3] (without 3)
}
```
