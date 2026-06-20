# 1063 — Number Of Valid Subarrays

## Deskripsi

**Soal:** [1063. Number Of Valid Subarrays](https://leetcode.com/problems/number-of-valid-subarrays/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO), Monotonic Stack (tumpukan monoton)

## Solusi Go

```go
package main

// LeetCode #1063: Number of Valid Subarrays
// https://leetcode.com/problems/number-of-valid-subarrays/
// Difficulty: Hard [Paid]
//
// Monotonic stack approach. A valid subarray is one where the first element
// is the minimum of that subarray. For each element at index i, we find the
// next smaller element to the right (at index j). All subarrays starting at
// i and ending before j have arr[i] as the minimum, so count += j-i.

import "fmt"

func main() {
	fmt.Println(validSubarrays([]int{1, 4, 2, 5, 3}))
}

func validSubarrays(nums []int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0)
	count := 0

	for i := 0; i <= n; i++ {
		// Pop elements while current value is smaller than top of stack
		for len(stack) > 0 && (i == n || nums[stack[len(stack)-1]] > nums[i]) {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			count += i - idx
		}
		if i < n {
			stack = append(stack, i)
		}
	}

	return count
}
```
