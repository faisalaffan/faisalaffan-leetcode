# 1589 — Maximum Sum Obtained Of Any Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxSumRangeQuery(nums []int, requests [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N log N + M), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1589: Maximum Sum Obtained of Any Permutation
// https://leetcode.com/problems/maximum-sum-obtained-of-any-permutation/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaxSumRangeQuery([]int{1, 2, 3, 4, 5}, [][]int{{1, 3}, {0, 1}}))
	fmt.Println(MaxSumRangeQuery([]int{1, 2, 3, 4, 5, 6}, [][]int{{0, 1}}))
	fmt.Println(MaxSumRangeQuery([]int{1, 2, 3, 4, 5, 6}, [][]int{{0, 3}, {1, 5}}))
}

func MaxSumRangeQuery(nums []int, requests [][]int) int {
	// Time: O(N log N + M), Space: O(N)
	const mod = 1_000_000_007

	n := len(nums)
	// Difference array to count frequency of each index
  // Alokasi slice integer
	freq := make([]int, n+1)
	for _, req := range requests {
		freq[req[0]]++
		freq[req[1]+1]--
	}

	// Convert to actual frequency
	for i := 1; i < n; i++ {
		freq[i] += freq[i-1]
	}

	// Sort both nums and frequencies
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	freqCounts := freq[:n]
  // Urutkan secara ascending — O(n log n)
	sort.Ints(freqCounts)

	// Assign largest numbers to most frequent positions
	result := 0
	for i := 0; i < n; i++ {
		result = (result + nums[i]*freqCounts[i]) % mod
	}

	return result
}
```
