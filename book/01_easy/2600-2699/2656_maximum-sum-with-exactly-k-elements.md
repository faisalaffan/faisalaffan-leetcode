# 2656 — Maximum Sum With Exactly K Elements

## Deskripsi

**Soal:** [2656. Maximum Sum With Exactly K Elements](https://leetcode.com/problems/maximum-sum-with-exactly-k-elements/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2656: Maximum Sum With Exactly K Elements
// https://leetcode.com/problems/maximum-sum-with-exactly-k-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaximumSumWithExactlyKElements([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(MaximumSumWithExactlyKElements([]int{5, 5, 5}, 2))
}

func MaximumSumWithExactlyKElements(nums []int, k int) int {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	// Sum of arithmetic series: maxVal + (maxVal+1) + ... + (maxVal+k-1)
	return k * (2*maxVal + k - 1) / 2
}
```
