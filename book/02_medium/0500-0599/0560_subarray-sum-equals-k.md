# 0560 — Subarray Sum Equals K

## Deskripsi

**Soal:** [0560. Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #560: Subarray Sum Equals K
// https://leetcode.com/problems/subarray-sum-equals-k/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(SubarraySum([]int{1, 1, 1}, 2))
	fmt.Println(SubarraySum([]int{1, 2, 3}, 3))
}

func SubarraySum(nums []int, k int) int {
	count := 0
	sum := 0
  // Membuat map untuk pencarian O(1): key → value
	sumMap := make(map[int]int)
	sumMap[0] = 1

	for _, num := range nums {
		sum += num
		if val, ok := sumMap[sum-k]; ok {
			count += val
		}
		sumMap[sum]++
	}

	return count
}
```
