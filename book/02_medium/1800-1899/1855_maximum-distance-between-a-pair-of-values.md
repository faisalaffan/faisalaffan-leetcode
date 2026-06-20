# 1855 — Maximum Distance Between A Pair Of Values

## Deskripsi

**Soal:** [1855. Maximum Distance Between A Pair Of Values](https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m+n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1855: Maximum Distance Between a Pair of Values
// https://leetcode.com/problems/maximum-distance-between-a-pair-of-values/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}))
	fmt.Println(MaxDistance([]int{2, 2, 2}, []int{10, 10, 1}))
	fmt.Println(MaxDistance([]int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}))
}

// Time: O(m+n), Space: O(1)
func MaxDistance(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	maxDist := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			if j-i > maxDist {
				maxDist = j - i
			}
			j++
		} else {
			i++
		}
	}
	return maxDist
}
```
