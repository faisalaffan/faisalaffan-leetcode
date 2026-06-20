# 2902 — Count Of Sub Multisets With Bounded Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func countSubMultisetsWithBoundedSum(nums []int, l, r int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Sliding Window, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2902: Count of Sub-Multisets With Bounded Sum
// https://leetcode.com/problems/count-of-sub-multisets-with-bounded-sum/
// Difficulty: Hard
//
// Count the number of sub-multisets (not necessarily contiguous) of nums whose sum
// is between l and r inclusive. Use bounded knapsack DP with sliding window
// optimization. For each value v with frequency f, process the DP array using
// a sliding window over each residue class modulo v to achieve O(total * unique)
// time where total = sum(nums) <= 10^5.
//
// n <= 10^5, sum(nums) <= 10^5.

import "fmt"

const mod2902 = 1000000007

func countSubMultisetsWithBoundedSum(nums []int, l, r int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	totalSum := 0
	for _, v := range nums {
		freq[v]++
		totalSum += v
	}

	maxSum := r
	if totalSum < maxSum {
		maxSum = totalSum
	}
	if maxSum < 0 {
		return 0
	}

  // Alokasi slice integer
	dp := make([]int, maxSum+1)
	dp[0] = 1

	for v, f := range freq {
  // Alokasi slice integer
		newdp := make([]int, maxSum+1)
		copy(newdp, dp)

		for rem := 0; rem < v && rem <= maxSum; rem++ {
			var window int64
			for pos := rem; pos <= maxSum; pos += v {
				// Add element entering the window
				window += int64(dp[pos])
				// Remove element leaving the window
				removePos := pos - (f+1)*v
				if removePos >= 0 {
					window -= int64(dp[removePos])
				}
				// window = sum of dp at positions pos, pos-v, ..., pos-f*v
				// This represents using 0, 1, ..., f copies of value v to reach sum pos
				modVal := int(window % mod2902)
				if modVal < 0 {
					modVal += mod2902
				}
				newdp[pos] = (newdp[pos] + modVal) % mod2902
			}
		}
		dp = newdp
	}

	ans := 0
	for s := l; s <= maxSum; s++ {
		ans = (ans + dp[s]) % mod2902
	}
	return ans
}

func main() {
	// Example: nums=[2,2,3], l=3, r=5
	// Sub-multisets: {3}(3), {2,2}(4), {2,3}(5) => 3
	fmt.Println(countSubMultisetsWithBoundedSum([]int{2, 2, 3}, 3, 5))

	// Example: nums=[1,2,3], l=1, r=3
	// {1}, {2}, {3}, {1,2}, {2,1 is same}... wait sub-multisets: {1}(1), {2}(2), {3}(3), {1,2}(3) => 4
	// Actually {1,2} has sum 3. So in range [1,3]: {1}, {2}, {3}, {1,2} => 4
	fmt.Println(countSubMultisetsWithBoundedSum([]int{1, 2, 3}, 1, 3))

	// All ones, sum=5, total sum = 5
	fmt.Println(countSubMultisetsWithBoundedSum([]int{1, 1, 1, 1, 1}, 2, 4))

	// Single element
	fmt.Println(countSubMultisetsWithBoundedSum([]int{5}, 5, 5))

	// No valid submultisets
	fmt.Println(countSubMultisetsWithBoundedSum([]int{10}, 1, 5))

	// l=0 means empty set counts
	fmt.Println(countSubMultisetsWithBoundedSum([]int{1, 2}, 0, 0))

	// Duplicate values
	fmt.Println(countSubMultisetsWithBoundedSum([]int{2, 2, 2, 2}, 2, 6))

	// Larger range
	fmt.Println(countSubMultisetsWithBoundedSum([]int{1, 1, 2, 3}, 0, 10))
}
```
