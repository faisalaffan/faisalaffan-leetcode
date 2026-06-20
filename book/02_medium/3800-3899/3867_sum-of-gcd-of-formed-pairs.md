# 3867 — Sum Of Gcd Of Formed Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfGcdOfFormedPairs(nums []int) int
```

> **💡 Hint:** Build prefixGcd array where prefixGcd[i] = gcd(nums[i], max(nums[0..i])).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum, GCD / Matematika

**Kompleksitas Waktu:** O(N log M)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3867: Sum of GCD of Formed Pairs
// https://leetcode.com/problems/sum-of-gcd-of-formed-pairs/
// Difficulty: Medium
// Time: O(N log M) | Space: O(N)
// Approach: Build prefixGcd array where prefixGcd[i] = gcd(nums[i], max(nums[0..i])).
// Sort, pair smallest with largest, sum gcd of each pair.

import (
	"fmt"
	"sort"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func SumOfGcdOfFormedPairs(nums []int) int {
	n := len(nums)
  // Alokasi slice integer
	prefixGcd := make([]int, n)
	mx := 0
	for i, v := range nums {
		if v > mx {
			mx = v
		}
		prefixGcd[i] = gcd(v, mx)
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(prefixGcd)

	ans := 0
	for i := 0; i < n/2; i++ {
		ans += gcd(prefixGcd[i], prefixGcd[n-1-i])
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(SumOfGcdOfFormedPairs([]int{2, 6, 4})) // Expected: 2

	// Example 2
	fmt.Println(SumOfGcdOfFormedPairs([]int{3, 6, 2, 8})) // Expected: 5

	// Extra
	fmt.Println(SumOfGcdOfFormedPairs([]int{1})) // Expected: 0
}
```
