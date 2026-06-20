# 1909 — Remove One Element To Make The Array Strictly Increasing

## Deskripsi

**Soal:** [1909. Remove One Element To Make The Array Strictly Increasing](https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1909: Remove One Element to Make the Array Strictly Increasing
// https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RemoveOneElementToMakeTheArrayStrictlyIncreasing([]int{1, 2, 10, 5, 7}))       // true
	fmt.Println(RemoveOneElementToMakeTheArrayStrictlyIncreasing([]int{2, 3, 1, 2}))           // false
	fmt.Println(RemoveOneElementToMakeTheArrayStrictlyIncreasing([]int{1, 1, 1}))              // false
}

// Time: O(n), Space: O(1)
func RemoveOneElementToMakeTheArrayStrictlyIncreasing(nums []int) bool {
	removed := false
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if removed {
				return false
			}
			removed = true
			// Try removing nums[i-1] or nums[i]
			if i-2 < 0 || nums[i] > nums[i-2] {
				// Removing nums[i-1] works
			} else if i+1 >= len(nums) || nums[i+1] > nums[i-1] {
				// Removing nums[i] works, skip it
				nums[i] = nums[i-1] // adjust for next comparison
			} else {
				return false
			}
		}
	}
	return true
}
```
