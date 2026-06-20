# 3671 — Sum Of Beautiful Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumOfBeautifulSubsequences(nums []int) int
```

> **💡 Hint:** For each element as the minimum, count subsequences

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3671: Sum of Beautiful Subsequences
// https://leetcode.com/problems/sum-of-beautiful-subsequences/
// Difficulty: Hard
//
// A subsequence is "beautiful" if the product of its length and
// the minimum element is maximized. Sum the values of all beautiful
// subsequences (where value = length * min).
//
// Approach: For each element as the minimum, count subsequences
// where this element is the minimum and compute their total value.
// Use DP tracking contribution of each element.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumOfBeautifulSubsequences([]int{1, 2, 3}))
	// Example 2
	fmt.Println(sumOfBeautifulSubsequences([]int{3, 1, 2}))
	// Edge: single element
	fmt.Println(sumOfBeautifulSubsequences([]int{5}))
}

const B_MOD = 1000000007

func sumOfBeautifulSubsequences(nums []int) int {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	total := int64(0)

	// For each element as the minimum of the subsequence
	for i, minVal := range nums {
		// Count subsequences where nums[i] is the minimum
		// Elements before i that are >= minVal can be included or not
		leftChoices := 0
		for j := 0; j < i; j++ {
			if nums[j] >= minVal {
				leftChoices++
			}
		}
		// Elements after i that are > minVal can be included or not
		rightChoices := 0
		for j := i + 1; j < n; j++ {
			if nums[j] > minVal {
				rightChoices++
			}
		}

		// Number of subsequences where nums[i] is the minimum
		// Each left/right choice can be independently selected
		subseqCount := pow2(leftChoices) * pow2(rightChoices) % B_MOD

		// For each such subsequence, the value contributed is
		// sum over all lengths where nums[i] is included
		// length = 1 (just nums[i]) + additional elements from left/right
		// Each subset of left+right gives length = 1 + k where k = size of subset
		// Sum of value = minVal * sum of (1 + k) for each subset
		// = minVal * (totalSubseq * 1 + sum_of_k_over_all_subsets)
		// sum_of_k_over_all_subsets = (left+right) * 2^(left+right-1) for left+right > 0

		extra := leftChoices + rightChoices
		var sumLen int64
		if extra == 0 {
			sumLen = 1
		} else {
			sumLen = int64(pow2(extra))                              // all subsets (each of length >= 1)
			sumLen = (sumLen + int64(extra)*int64(pow2(extra-1))) % B_MOD // total additional elements across all subsets
		}

		contrib := int64(minVal) * sumLen % B_MOD
		contrib = contrib * int64(subseqCount) % B_MOD
		contrib = contrib * int64(pow2(int(B_MOD-2))) % B_MOD // Wrong approach, let me simplify

		// Simpler: total value = minVal * sum of lengths over all valid subsets
		// Each valid subset length = 1 + k where k elements from left+right
		// Count subsets of size k from extra elements
		totalValue := int64(0)
		for k := 0; k <= extra; k++ {
			ways := nCr(extra, k)
			length := 1 + k
			totalValue = (totalValue + int64(length)*int64(ways)%B_MOD) % B_MOD
		}
		contrib = int64(minVal) * totalValue % B_MOD
		contrib = contrib * int64(subseqCount) % B_MOD

		total = (total + contrib) % B_MOD
	}

	return int(total)
}

var fact []int64
var invFact []int64

func initFact(n int) {
	fact = make([]int64, n+1)
	invFact = make([]int64, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * int64(i) % B_MOD
	}
	invFact[n] = modPow(fact[n], B_MOD-2)
	for i := n - 1; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % B_MOD
	}
}

func nCr(n, r int) int64 {
	if r < 0 || r > n {
		return 0
	}
	if fact == nil || len(fact) <= n {
		initFact(n + 10)
	}
	return fact[n] * invFact[r] % B_MOD * invFact[n-r] % B_MOD
}

func pow2(e int) int64 {
	return modPow(2, e)
}

func modPow(a int64, b int) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = res * a % B_MOD
		}
		a = a * a % B_MOD
		b >>= 1
	}
	return res
}

// Keep compiler happy
var _ = fmt.Println
```
