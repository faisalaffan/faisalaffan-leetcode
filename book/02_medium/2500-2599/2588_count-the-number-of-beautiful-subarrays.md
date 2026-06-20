# 2588 — Count The Number Of Beautiful Subarrays

## Deskripsi

**Soal:** [2588. Count The Number Of Beautiful Subarrays](https://leetcode.com/problems/count-the-number-of-beautiful-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func beautifulSubarrays(nums []int) int64`

## Solusi Go

```go
package main

// LeetCode #2588: Count the Number of Beautiful Subarrays
// https://leetcode.com/problems/count-the-number-of-beautiful-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func beautifulSubarrays(nums []int) int64 {
  // Membuat map untuk pencarian O(1): key → value
	prefixXor := make(map[int]int)
	prefixXor[0] = 1
	xor := 0
	var ans int64

	for _, v := range nums {
		xor ^= v
		ans += int64(prefixXor[xor])
		prefixXor[xor]++
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", beautifulSubarrays([]int{4, 3, 1, 2, 4}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", beautifulSubarrays([]int{1, 10, 4}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", beautifulSubarrays([]int{0, 0, 0}))
	// Expected: 6
}
```
