# 0667 — Beautiful Arrangement Ii

## Deskripsi

**Soal:** [0667. Beautiful Arrangement Ii](https://leetcode.com/problems/beautiful-arrangement-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #667: Beautiful Arrangement II
// https://leetcode.com/problems/beautiful-arrangement-ii/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(constructArray(3, 1))
	fmt.Println(constructArray(3, 2))
}

func constructArray(n int, k int) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
	left, right := 1, k+1

	for i := 0; i <= k; i++ {
		if i%2 == 0 {
			result[i] = left
			left++
		} else {
			result[i] = right
			right--
		}
	}

	for i := k + 1; i < n; i++ {
		result[i] = i + 1
	}

	return result
}
```
