# 0697 — Degree Of An Array

## Deskripsi

**Soal:** [0697. Degree Of An Array](https://leetcode.com/problems/degree-of-an-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #697: Degree of an Array
// https://leetcode.com/problems/degree-of-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))       // 2
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2})) // 6
	fmt.Println(findShortestSubArray([]int{1}))                    // 1
}

// findShortestSubArray finds the smallest subarray length with the same degree as the array.
// Time: O(n). Space: O(n).
func findShortestSubArray(nums []int) int {
  // Membuat map untuk pencarian O(1): key → value
	first := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	count := make(map[int]int)
	maxCount := 0
	minLen := len(nums)

	for i, v := range nums {
		if _, ok := first[v]; !ok {
			first[v] = i
		}
		count[v]++
		if count[v] > maxCount {
			maxCount = count[v]
		}
	}

	for v, c := range count {
		if c == maxCount {
			// find last occurrence
			last := 0
			for i := len(nums) - 1; i >= 0; i-- {
				if nums[i] == v {
					last = i
					break
				}
			}
			length := last - first[v] + 1
			if length < minLen {
				minLen = length
			}
		}
	}
	return minLen
}
```
