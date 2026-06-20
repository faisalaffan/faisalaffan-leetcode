# 3915 — Maximum Sum Of Alternating Subsequence With Distance At Least K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxAlternatingSum(nums []int, k int) int64
```

> **💡 Hint:** DP. For each position i, track max alternating sum

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3915: Maximum Sum of Alternating Subsequence With Distance at Least K
// https://leetcode.com/problems/maximum-sum-of-alternating-subsequence-with-distance-at-least-k/
// Difficulty: Hard
//
// Pick a subsequence where adjacent selected elements have index
// distance at least k. Maximize alternating sum: elements at even
// positions (0-indexed in subsequence) are added, odd positions
// are subtracted.
//
// Approach: DP. For each position i, track max alternating sum
// ending at i with even length and odd length.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxAlternatingSum([]int{1, 2, 3, 4, 5}, 2))
	// Example 2
	fmt.Println(maxAlternatingSum([]int{10, 1, 10, 1, 10}, 2))
	// Edge: k = 1
	fmt.Println(maxAlternatingSum([]int{1, 2, 3}, 1))
}

func maxAlternatingSum(nums []int, k int) int64 {
	n := len(nums)
	// dpEven[i] = max alternating sum ending at i with even length (positive)
	// dpOdd[i] = max alternating sum ending at i with odd length (negative end)
  // Alokasi slice integer
	dpEven := make([]int64, n)
  // Alokasi slice integer
	dpOdd := make([]int64, n)

	for i := 0; i < n; i++ {
		// Start new subsequence at i
		dpEven[i] = int64(nums[i])
		dpOdd[i] = -int64(1 << 60) // impossible

		// Try extending from earlier positions
		for j := 0; j <= i-k; j++ {
			// Even length: extend from odd (add nums[i])
			if dpOdd[j]+int64(nums[i]) > dpEven[i] {
				dpEven[i] = dpOdd[j] + int64(nums[i])
			}
			// Odd length: extend from even (subtract nums[i])
			if dpEven[j]-int64(nums[i]) > dpOdd[i] {
				dpOdd[i] = dpEven[j] - int64(nums[i])
			}
		}
	}

	var ans int64
	for i := 0; i < n; i++ {
		if dpEven[i] > ans {
			ans = dpEven[i]
		}
	}
	return ans
}
```
