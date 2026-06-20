# 3315 — Construct The Minimum Bitwise Array Ii

## Deskripsi

**Soal:** [3315. Construct The Minimum Bitwise Array Ii](https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log m) Space: O(1) (excluding output)  
**Kompleksitas Ruang:** O(1) (excluding output)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3315: Construct the Minimum Bitwise Array II
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/
// Difficulty: Medium
// Time: O(n log m) Space: O(1) (excluding output)

import "fmt"

func main() {
	fmt.Println(minBitwiseArray([]int{11, 13, 31})) // [9 12 15]
	fmt.Println(minBitwiseArray([]int{2, 3, 5}))    // [-1 1 4]
	fmt.Println(minBitwiseArray([]int{7}))           // [3]
}

func minBitwiseArray(nums []int) []int {
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(nums))
	for i, num := range nums {
		if num == 2 {
			ans[i] = -1
			continue
		}
		// Find rightmost block of 1s in binary
		p := 0
		for (num>>p)&1 == 1 {
			p++
		}
		ans[i] = num ^ (1 << (p - 1))
	}
	return ans
}
```
