# 2831 — Find The Longest Equal Subarray

## Deskripsi

**Soal:** [2831. Find The Longest Equal Subarray](https://leetcode.com/problems/find-the-longest-equal-subarray/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func FindTheLongestEqualSubarray(nums []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #2831: Find the Longest Equal Subarray
// https://leetcode.com/problems/find-the-longest-equal-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func FindTheLongestEqualSubarray(nums []int, k int) int {
  // Membuat map untuk pencarian O(1): key → value
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	best := 0
	for _, indices := range pos {
		left := 0
		for right := 0; right < len(indices); right++ {
			// Elements between indices[left] and indices[right] that need to be removed
			for indices[right]-indices[left]-(right-left) > k {
				left++
			}
			if right-left+1 > best {
				best = right - left + 1
			}
		}
	}

	return best
}

func main() {
	fmt.Println(FindTheLongestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Println(FindTheLongestEqualSubarray([]int{1, 1, 2, 2, 1, 1}, 2))
}
```
