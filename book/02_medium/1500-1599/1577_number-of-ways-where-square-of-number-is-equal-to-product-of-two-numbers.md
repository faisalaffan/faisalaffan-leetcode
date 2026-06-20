# 1577 — Number Of Ways Where Square Of Number Is Equal To Product Of Two Numbers

## Deskripsi

**Soal:** [1577. Number Of Ways Where Square Of Number Is Equal To Product Of Two Numbers](https://leetcode.com/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N^2 + M^2), Space: O(N^2)  
**Kompleksitas Ruang:** O(N^2)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1577: Number of Ways Where Square of Number Is Equal to Product of Two Numbers
// https://leetcode.com/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumTriplets([]int{7, 4}, []int{5, 2, 8, 9}))
	fmt.Println(NumTriplets([]int{1, 1}, []int{1, 1, 1}))
	fmt.Println(NumTriplets([]int{7, 7, 8, 3}, []int{1, 2, 9, 7}))
}

func NumTriplets(nums1 []int, nums2 []int) int {
	// Time: O(N^2 + M^2), Space: O(N^2)
	// Count pairs in each array that multiply to a specific product
	return countSquareProducts(nums1, nums2) + countSquareProducts(nums2, nums1)
}

func countSquareProducts(nums1 []int, nums2 []int) int {
	// Count nums1[i]^2 == nums2[j] * nums2[k] for j < k
  // Membuat map untuk pencarian O(1): key → value
	productCount := make(map[int]int)
	for j := 0; j < len(nums2); j++ {
		for k := j + 1; k < len(nums2); k++ {
			product := nums2[j] * nums2[k]
			productCount[product]++
		}
	}

	count := 0
	for _, v := range nums1 {
		square := v * v
		count += productCount[square]
	}

	return count
}
```
