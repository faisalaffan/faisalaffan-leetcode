# 3489 — Zero Array Transformation Iv

## Deskripsi

**Soal:** [3489. Zero Array Transformation Iv](https://leetcode.com/problems/zero-array-transformation-iv/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3489: Zero Array Transformation IV
// https://leetcode.com/problems/zero-array-transformation-iv/
// Difficulty: Medium
// Complexity: O(n * q * log n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	nums := []int{3, 5, 2}
	queries := [][]int{{0, 2, 1}, {1, 3, 2}, {0, 3, 3}}
	fmt.Println("Test 1:", ZeroArrayTransformationIv(nums, queries))

	// Test case 2
	nums2 := []int{1, 2, 3}
	queries2 := [][]int{{0, 1, 1}, {1, 2, 2}}
	fmt.Println("Test 2:", ZeroArrayTransformationIv(nums2, queries2))

	// Test case 3
	nums3 := []int{0, 0, 0}
	queries3 := [][]int{{0, 2, 1}}
	fmt.Println("Test 3:", ZeroArrayTransformationIv(nums3, queries3))
}

func ZeroArrayTransformationIv(nums []int, queries [][]int) int {
	// For each element, find the minimum number of queries needed to reduce it to 0
	// queries[i] = [l, r, val] meaning we can subtract val in range [l, r]
	minQueries := -1
	left, right := 0, len(queries)
	for left <= right {
		mid := left + (right-left)/2
		if canTransform(nums, queries, mid) {
			minQueries = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return minQueries
}

func canTransform(nums []int, queries [][]int, k int) bool {
	n := len(nums)
	// diff array to apply range updates
  // Membuat slice untuk menyimpan hasil
	diff := make([]int, n+1)
	for i := 0; i < k && i < len(queries); i++ {
		l, r, val := queries[i][0], queries[i][1], queries[i][2]
		if l >= n {
			continue
		}
		if r >= n-1 {
			r = n - 1
		}
		diff[l] += val
		diff[r+1] -= val
	}

	cur := 0
	for i := 0; i < n; i++ {
		cur += diff[i]
		if cur < nums[i] {
			return false
		}
	}
	return true
}
```
