# 0377 — Combination Sum Iv

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func combinationSum4(nums []int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(target * n)  
**Kompleksitas Ruang:** O(target)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #377: Combination Sum IV
// https://leetcode.com/problems/combination-sum-iv/
// Difficulty: Medium
// Time: O(target * n) | Space: O(target)

import "fmt"

func combinationSum4(nums []int, target int) int {
  // Alokasi slice integer
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", combinationSum4([]int{1, 2, 3}, 4))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", combinationSum4([]int{9}, 3))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", combinationSum4([]int{1, 2, 3}, 3))
	// Expected: 4
}
```
