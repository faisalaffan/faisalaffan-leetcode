# 1966 — Binary Searchable Numbers In An Unsorted Array

## Deskripsi

**Soal:** [1966. Binary Searchable Numbers In An Unsorted Array](https://leetcode.com/problems/binary-searchable-numbers-in-an-unsorted-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Binary Search (pencarian biner)

## Solusi Go

```go
package main

// LeetCode #1966: Binary Searchable Numbers in an Unsorted Array
// https://leetcode.com/problems/binary-searchable-numbers-in-an-unsorted-array/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(BinarySearchableNumbers([]int{2, 1, 3, 5, 4, 6}))
	fmt.Println(BinarySearchableNumbers([]int{1, 3, 2}))
	fmt.Println(BinarySearchableNumbers([]int{2, 3, 1}))
}

// Time: O(n), Space: O(n)
func BinarySearchableNumbers(nums []int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	prefixMax := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	suffixMin := make([]int, n)

	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefixMax[i-1] {
			prefixMax[i] = nums[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}

	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if prefixMax[i] == suffixMin[i] {
			count++
		}
	}
	return count
}
```
