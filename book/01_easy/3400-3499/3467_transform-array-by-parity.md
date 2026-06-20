# 3467 — Transform Array By Parity

## Deskripsi

**Soal:** [3467. Transform Array By Parity](https://leetcode.com/problems/transform-array-by-parity/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3467: Transform Array by Parity
// https://leetcode.com/problems/transform-array-by-parity/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TransformArrayByParity([]int{4, 3, 2, 1}))
	fmt.Println(TransformArrayByParity([]int{1, 5, 2, 8, 3}))
}

// TransformArrayByParity transforms array: even numbers -> 0 (sorted first), odd numbers -> 1.
// Time: O(n log n). Space: O(1).
func TransformArrayByParity(nums []int) []int {
	// Count evens
	evenCount := 0
	for _, v := range nums {
		if v%2 == 0 {
			evenCount++
		}
	}
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(nums))
	for i := 0; i < evenCount; i++ {
		result[i] = 0
	}
	for i := evenCount; i < len(nums); i++ {
		result[i] = 1
	}
	return result
}
```
