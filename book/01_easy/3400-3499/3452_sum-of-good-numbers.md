# 3452 — Sum Of Good Numbers

## Deskripsi

**Soal:** [3452. Sum Of Good Numbers](https://leetcode.com/problems/sum-of-good-numbers/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3452: Sum of Good Numbers
// https://leetcode.com/problems/sum-of-good-numbers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfGoodNumbers([]int{1, 3, 2, 1, 5, 4}, 2))
	fmt.Println(SumOfGoodNumbers([]int{2, 1}, 1))
}

// SumOfGoodNumbers returns sum of numbers that are greater than both nums[i-k] and nums[i+k] (or if out of bounds).
// Time: O(n). Space: O(1).
func SumOfGoodNumbers(nums []int, k int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		good := true
		if i-k >= 0 && nums[i] <= nums[i-k] {
			good = false
		}
		if i+k < n && nums[i] <= nums[i+k] {
			good = false
		}
		if good {
			sum += nums[i]
		}
	}
	return sum
}
```
