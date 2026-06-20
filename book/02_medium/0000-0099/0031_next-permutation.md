# 0031 — Next Permutation

## Deskripsi

**Soal:** [0031. Next Permutation](https://leetcode.com/problems/next-permutation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func nextPermutation(nums []int) `

## Solusi Go

```go
package main

// LeetCode #31: Next Permutation
// https://leetcode.com/problems/next-permutation/
// Difficulty: Medium

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// Find first decreasing element from right
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Reverse suffix
	for left, right := i+1, n-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

func main() {
	// Test case 1
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // [1 3 2]

	// Test case 2
	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums) // [1 2 3]

	// Test case 3
	nums = []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Println(nums) // [1 5 1]
}

// Time: O(n) | Space: O(1)
```
