# 2239 — Find Closest Number To Zero

## Deskripsi

**Soal:** [2239. Find Closest Number To Zero](https://leetcode.com/problems/find-closest-number-to-zero/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2239: Find Closest Number to Zero
// https://leetcode.com/problems/find-closest-number-to-zero/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindClosestNumberToZero([]int{-4, -2, 1, 4, 8})) // 1
	fmt.Println(FindClosestNumberToZero([]int{2, -1, 1}))         // 1
}

// Time: O(n), Space: O(1)
func FindClosestNumberToZero(nums []int) int {
	closest := nums[0]
	for _, v := range nums[1:] {
		absV := v
		if absV < 0 {
			absV = -absV
		}
		absClosest := closest
		if absClosest < 0 {
			absClosest = -absClosest
		}
		if absV < absClosest || (absV == absClosest && v > closest) {
			closest = v
		}
	}
	return closest
}
```
