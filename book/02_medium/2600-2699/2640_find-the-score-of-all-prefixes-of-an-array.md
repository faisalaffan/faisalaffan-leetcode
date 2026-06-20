# 2640 — Find The Score Of All Prefixes Of An Array

## Deskripsi

**Soal:** [2640. Find The Score Of All Prefixes Of An Array](https://leetcode.com/problems/find-the-score-of-all-prefixes-of-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func findPrefixScore(nums []int) []int64`

## Solusi Go

```go
package main

// LeetCode #2640: Find the Score of All Prefixes of an Array
// https://leetcode.com/problems/find-the-score-of-all-prefixes-of-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findPrefixScore(nums []int) []int64 {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int64, n)
	maxSoFar := nums[0]
	prefixSum := int64(0)

	for i, v := range nums {
		if v > maxSoFar {
			maxSoFar = v
		}
		conver := int64(v + maxSoFar)
		prefixSum += conver
		ans[i] = prefixSum
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findPrefixScore([]int{2, 3, 7, 5, 10}))
	// Expected: [4,10,24,36,56]

	// Test case 2
	fmt.Println("Test 2:", findPrefixScore([]int{1, 1, 1}))
	// Expected: [2,4,6]

	// Test case 3
	fmt.Println("Test 3:", findPrefixScore([]int{5}))
	// Expected: [10]
}
```
