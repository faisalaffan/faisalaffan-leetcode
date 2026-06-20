# 2970 — Count The Number Of Incremovable Subarrays I

## Deskripsi

**Soal:** [2970. Count The Number Of Incremovable Subarrays I](https://leetcode.com/problems/count-the-number-of-incremovable-subarrays-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^3)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2970: Count the Number of Incremovable Subarrays I
// https://leetcode.com/problems/count-the-number-of-incremovable-subarrays-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: incremovableSubarrayCount
	fmt.Println(CountTheNumberOfIncremovableSubarraysI([]int{1, 2, 3, 4})) // 10
	fmt.Println(CountTheNumberOfIncremovableSubarraysI([]int{6, 5, 7, 8})) // 7
	fmt.Println(CountTheNumberOfIncremovableSubarraysI([]int{8, 7, 6, 6})) // 3
}

// Time: O(n^3) | Space: O(n)
// LeetCode submission name: incremovableSubarrayCount
func CountTheNumberOfIncremovableSubarraysI(nums []int) int {
	n := len(nums)
	count := 0

	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			// Check if array without nums[l..r] is strictly increasing
			if isStrictlyIncreasing(nums, l, r) {
				count++
			}
		}
	}
	return count
}

func isStrictlyIncreasing(nums []int, l, r int) bool {
	prev := -1
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		if i >= l && i <= r {
			continue
		}
		if nums[i] <= prev {
			return false
		}
		prev = nums[i]
	}
	return true
}
```
