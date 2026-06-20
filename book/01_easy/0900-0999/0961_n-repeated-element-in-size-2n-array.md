# 0961 — N Repeated Element In Size 2N Array

## Deskripsi

**Soal:** [0961. N Repeated Element In Size 2N Array](https://leetcode.com/problems/n-repeated-element-in-size-2n-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #961: N-Repeated Element in Size 2N Array
// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3})) // 3
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2})) // 2
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4})) // 5
}

// repeatedNTimes finds the element repeated n times in a 2n size array.
// Time: O(n). Space: O(1).
func repeatedNTimes(nums []int) int {
	// Since the element appears n times in 2n, any two consecutive elements
	// must contain the repeated element (in most cases).
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == nums[i+1] || nums[i] == nums[i+2] {
			return nums[i]
		}
	}
	// If not found yet, the repeated element is in the last 3 positions
	return nums[len(nums)-1]
}
```
