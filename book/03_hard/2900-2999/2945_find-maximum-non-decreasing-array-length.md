# 2945 — Find Maximum Non Decreasing Array Length

## Deskripsi

**Soal:** [2945. Find Maximum Non Decreasing Array Length](https://leetcode.com/problems/find-maximum-non-decreasing-array-length/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Prefix Sum (jumlah kumulatif)

**Fungsi Solusi:** `func findMaximumLength(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2945: Find Maximum Non-decreasing Array Length
// https://leetcode.com/problems/find-maximum-non-decreasing-array-length/
// Difficulty: Hard
//
// DP + monotonic deque optimization.
// Partition array into contiguous groups, replace each group with its sum.
// Goal: resulting array is non-decreasing, maximize number of groups.
//
// Define:
//   f[i] = max groups for prefix ending at i-1 (i elements)
//   g[i] = minimum possible last group sum achieving f[i]
//   pref[i] = prefix sum of first i elements
//
// Transition: for j < i where sum(j..i-1) = pref[i]-pref[j] >= g[j],
//   f[i] = f[j] + 1, g[i] = pref[i]-pref[j]
//
// Optimization: deque maintains candidates sorted by g[j] and pref[j]+g[j].

import (
	"fmt"
)

func findMaximumLength(nums []int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	pref := make([]int64, n+1)
	for i, v := range nums {
		pref[i+1] = pref[i] + int64(v)
	}

  // Membuat slice untuk menyimpan hasil
	f := make([]int, n+1)
  // Membuat slice untuk menyimpan hasil
	g := make([]int64, n+1)
  // Membuat slice untuk menyimpan hasil
	deq := make([]int, 0, n+1)
	deq = append(deq, 0)

	head := 0
	for i := 1; i <= n; i++ {
		// Pop front: discard indices that are no longer optimal
		for head+1 < len(deq) && pref[i] >= pref[deq[head+1]]+g[deq[head+1]] {
			head++
		}

		j := deq[head]
		f[i] = f[j] + 1
		g[i] = pref[i] - pref[j]

		// Pop back: maintain monotonicity of pref[i]+g[i]
		for len(deq) > head && pref[i]+g[i] <= pref[deq[len(deq)-1]]+g[deq[len(deq)-1]] {
			deq = deq[:len(deq)-1]
		}
		deq = append(deq, i)
	}
	return f[n]
}

func main() {
	// Example: [2,3,1,4,5] -> 4
	// Partition: [2],[3],[1,4],[5] -> sums [2,3,5,5] (non-decreasing, length 4)
	fmt.Println(findMaximumLength([]int{2, 3, 1, 4, 5}))

	// Simple cases
	fmt.Println(findMaximumLength([]int{1, 2, 3}))
	fmt.Println(findMaximumLength([]int{5, 4, 3, 2, 1}))
	fmt.Println(findMaximumLength([]int{1, 1, 1}))
	fmt.Println(findMaximumLength([]int{1, 2, 1, 2}))
}
```
