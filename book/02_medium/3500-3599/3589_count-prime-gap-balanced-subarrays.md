# 3589 — Count Prime Gap Balanced Subarrays

## Deskripsi

**Soal:** [3589. Count Prime Gap Balanced Subarrays](https://leetcode.com/problems/count-prime-gap-balanced-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func isPrime(x int) bool`

## Solusi Go

```go
package main

// LeetCode #3589: Count Prime-Gap Balanced Subarrays
// https://leetcode.com/problems/count-prime-gap-balanced-subarrays/
// Difficulty: Medium
// Complexity: O(n * sqrt(m)) time, O(n) space

import "fmt"

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", CountPrimeGapBalancedSubarrays([]int{2, 3, 5, 7}))
	// Test case 2
	fmt.Println("Test 2:", CountPrimeGapBalancedSubarrays([]int{4, 6, 8, 10}))
	// Test case 3
	fmt.Println("Test 3:", CountPrimeGapBalancedSubarrays([]int{2, 4, 6, 3}))
}

func CountPrimeGapBalancedSubarrays(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		primeCount := 0
		for j := i; j < n; j++ {
			if isPrime(nums[j]) {
				primeCount++
			}
			gap := j - i + 1
			nonPrimeCount := gap - primeCount
			if primeCount == nonPrimeCount {
				count++
			}
		}
	}
	return count
}
```
