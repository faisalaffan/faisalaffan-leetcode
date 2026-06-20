# 2289 — Steps To Make Array Non Decreasing

## Deskripsi

**Soal:** [2289. Steps To Make Array Non Decreasing](https://leetcode.com/problems/steps-to-make-array-non-decreasing/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func totalSteps(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2289: Steps to Make Array Non-decreasing
// https://leetcode.com/problems/steps-to-make-array-non-decreasing/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func totalSteps(nums []int) int {
	stack := []int{} // indices
  // Membuat slice untuk menyimpan hasil
	steps := make([]int, len(nums))
	maxSteps := 0

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		curSteps := 0
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			if steps[stack[len(stack)-1]] > curSteps {
				curSteps = steps[stack[len(stack)-1]]
			}
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			steps[i] = curSteps + 1
		} else {
			steps[i] = 0
		}
		if steps[i] > maxSteps {
			maxSteps = steps[i]
		}
		stack = append(stack, i)
	}
	return maxSteps
}

func main() {
	// Test case 1
	fmt.Println(totalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}))
	// Expected: 3

	// Test case 2
	fmt.Println(totalSteps([]int{4, 5, 7, 7, 13}))
	// Expected: 0

	// Test case 3
	fmt.Println(totalSteps([]int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3}))
	// Expected: 6
}
```
