# 2453 — Destroy Sequential Targets

## Deskripsi

**Soal:** [2453. Destroy Sequential Targets](https://leetcode.com/problems/destroy-sequential-targets/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2453: Destroy Sequential Targets
// https://leetcode.com/problems/destroy-sequential-targets/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Group nums by nums[i] % space. The group with max size gives max targets.
// Pick smallest nums[i] from that group.

import "fmt"

func main() {
	fmt.Println(destroyTargets([]int{3, 7, 8, 1, 1, 5}, 2)) // 1
	fmt.Println(destroyTargets([]int{1, 3, 5, 2, 4, 6}, 2)) // 1
}

func destroyTargets(nums []int, space int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	minVal := make(map[int]int)

	for _, v := range nums {
		rem := v % space
		freq[rem]++
		if _, ok := minVal[rem]; !ok || v < minVal[rem] {
			minVal[rem] = v
		}
	}

	maxFreq, ans := 0, 0
	for rem, f := range freq {
		if f > maxFreq || (f == maxFreq && minVal[rem] < ans) {
			maxFreq = f
			ans = minVal[rem]
		}
	}
	return ans
}
```
