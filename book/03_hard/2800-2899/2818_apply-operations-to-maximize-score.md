# 2818 — Apply Operations To Maximize Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func primeScore(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting, Stack, Monotonic Stack

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2818: Apply Operations to Maximize Score
// https://leetcode.com/problems/apply-operations-to-maximize-score/
// Difficulty: Hard
//
// Monotonic stack to find subarray dominance + prime score + greedy. For each
// element compute its prime score (distinct prime factors), then find how many
// subarrays it dominates. Sort by value descending and apply k operations.
// O(N * sqrt(maxVal) + N log N) time, O(N) space.

import (
	"fmt"
	"sort"
)

const mod2818 = 1000000007

func primeScore(n int) int {
	count := 0
	remaining := n
	for p := 2; p*p <= remaining; p++ {
		if remaining%p == 0 {
			count++
			for remaining%p == 0 {
				remaining /= p
			}
		}
	}
	if remaining > 1 {
		count++
	}
	return count
}

func powMod(a, e int64) int64 {
	res := int64(1)
	a %= mod2818
	for e > 0 {
		if e&1 == 1 {
			res = (res * a) % mod2818
		}
		a = (a * a) % mod2818
		e >>= 1
	}
	return res
}

func maximumScore(nums []int, k int) int {
	n := len(nums)
  // Alokasi slice
	scores := make([]int, n)
	for i, v := range nums {
		scores[i] = primeScore(v)
	}

	// Previous greater (or equal) element index
  // Alokasi slice
	prev := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && scores[stack[len(stack)-1]] < scores[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			prev[i] = -1
		} else {
			prev[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Next greater (strictly greater) element index
  // Alokasi slice
	next := make([]int, n)
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && scores[stack[len(stack)-1]] <= scores[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			next[i] = n
		} else {
			next[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// Sort indices by value descending (if tie, by index ascending)
  // Alokasi slice
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
  // Custom sort
	sort.Slice(indices, func(i, j int) bool {
		if nums[indices[i]] != nums[indices[j]] {
			return nums[indices[i]] > nums[indices[j]]
		}
		return indices[i] < indices[j]
	})

	result := int64(1)
	for _, idx := range indices {
		leftCount := idx - prev[idx]
		rightCount := next[idx] - idx
		applications := leftCount * rightCount
		use := applications
		if use > k {
			use = k
		}
		if use > 0 {
			result = (result * powMod(int64(nums[idx]), int64(use))) % mod2818
			k -= use
		}
		if k == 0 {
			break
		}
	}

	return int(result)
}

func main() {
	// Example: nums=[8,3,9,3,8], k=2 => 81
	fmt.Println(maximumScore([]int{8, 3, 9, 3, 8}, 2))

	// Single element
	fmt.Println(maximumScore([]int{5}, 1))

	// k larger than total subarrays
	fmt.Println(maximumScore([]int{2, 3}, 3))

	// All same values
	fmt.Println(maximumScore([]int{4, 4, 4}, 2))

	// Prime-heavy inputs
	fmt.Println(maximumScore([]int{19, 12, 14, 6, 10}, 3))

	// Two elements
	fmt.Println(maximumScore([]int{2, 3}, 1))

	// k=1 with different values
	fmt.Println(maximumScore([]int{10, 20}, 1))

	// All prime numbers (score = 1)
	fmt.Println(maximumScore([]int{2, 3, 5, 7, 11, 13}, 4))

	// Large k
	fmt.Println(maximumScore([]int{100, 200, 300}, 10))

	// Values with many prime factors
	fmt.Println(maximumScore([]int{30, 42, 70, 105}, 5)) // 30=2*3*5, 42=2*3*7, 70=2*5*7, 105=3*5*7

	// Minimal case
	fmt.Println(maximumScore([]int{1}, 1)) // 1 has 0 prime factors
}
```
