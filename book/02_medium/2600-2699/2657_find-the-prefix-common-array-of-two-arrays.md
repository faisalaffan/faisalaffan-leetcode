# 2657 — Find The Prefix Common Array Of Two Arrays

## Deskripsi

**Soal:** [2657. Find The Prefix Common Array Of Two Arrays](https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func findThePrefixCommonArray(A []int, B []int) []int`

## Solusi Go

```go
package main

// LeetCode #2657: Find the Prefix Common Array of Two Arrays
// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)
	count := 0

	for i := 0; i < n; i++ {
		if seen[A[i]] {
			count++
		} else {
			seen[A[i]] = true
		}
		if seen[B[i]] {
			count++
		} else {
			seen[B[i]] = true
		}
		ans[i] = count
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
	// Expected: [0,2,3,4]

	// Test case 2
	fmt.Println("Test 2:", findThePrefixCommonArray([]int{1, 2, 3}, []int{1, 2, 3}))
	// Expected: [1,2,3]

	// Test case 3
	fmt.Println("Test 3:", findThePrefixCommonArray([]int{1, 2, 3}, []int{3, 1, 2}))
	// Expected: [0,1,3]
}
```
