# 0321 — Create Maximum Number

## Deskripsi

**Soal:** [0321. Create Maximum Number](https://leetcode.com/problems/create-maximum-number/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func maxNumber(nums1 []int, nums2 []int, k int) []int`

## Solusi Go

```go
package main

// LeetCode #321: Create Maximum Number
// https://leetcode.com/problems/create-maximum-number/
// Difficulty: Hard

import "fmt"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)

	// Pick max subsequence of length length from nums
	maxSubseq := func(nums []int, length int) []int {
		if length == 0 {
			return []int{}
		}
  // Membuat slice untuk menyimpan hasil
		stack := make([]int, 0, length)
		drop := len(nums) - length
		for _, v := range nums {
			for drop > 0 && len(stack) > 0 && stack[len(stack)-1] < v {
				stack = stack[:len(stack)-1]
				drop--
			}
			stack = append(stack, v)
		}
		return stack[:length]
	}

	// Merge two subsequences into lexicographically largest
	greater := func(a, b []int, i, j int) bool {
		for i < len(a) && j < len(b) {
			if a[i] != b[j] {
				return a[i] > b[j]
			}
			i++
			j++
		}
		return i < len(a)
	}

	merge := func(a, b []int) []int {
  // Membuat slice untuk menyimpan hasil
		res := make([]int, 0, len(a)+len(b))
		i, j := 0, 0
		for i < len(a) || j < len(b) {
			if j >= len(b) || (i < len(a) && greater(a, b, i, j)) {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
		return res
	}

	best := []int{}
	// Try all valid splits: take i from nums1, k-i from nums2
	start := 0
	if k > n {
		start = k - n
	}
	end := k
	if k > m {
		end = m
	}
	for i := start; i <= end; i++ {
		sub1 := maxSubseq(nums1, i)
		sub2 := maxSubseq(nums2, k-i)
		candidate := merge(sub1, sub2)
		if best == nil || greater(candidate, best, 0, 0) {
			best = candidate
		}
	}
	return best
}

func main() {
	// Example 1
	fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	// [9, 8, 6, 5, 3]

	// Example 2
	fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
	// [6, 7, 6, 0, 4]

	// Example 3
	fmt.Println(maxNumber([]int{3, 9}, []int{8, 9}, 3))
	// [9, 8, 9]
}
```
