# 0503 — Next Greater Element Ii

## Deskripsi

**Soal:** [0503. Next Greater Element Ii](https://leetcode.com/problems/next-greater-element-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

## Solusi Go

```go
package main

// LeetCode #503: Next Greater Element II
// https://leetcode.com/problems/next-greater-element-ii/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(NextGreaterElementIi([]int{1, 2, 1}))
	fmt.Println(NextGreaterElementIi([]int{1, 2, 3, 4, 3}))
}

func NextGreaterElementIi(nums []int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
  // Iterasi seluruh elemen
	for i := range result {
		result[i] = -1
	}

	stack := []int{}
	// Iterate twice to handle circular array
	for i := 0; i < 2*n; i++ {
		num := nums[i%n]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = num
		}
		if i < n {
			stack = append(stack, i)
		}
	}

	return result
}
```
