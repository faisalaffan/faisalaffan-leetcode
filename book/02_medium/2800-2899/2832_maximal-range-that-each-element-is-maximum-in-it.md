# 2832 — Maximal Range That Each Element Is Maximum In It

## Deskripsi

**Soal:** [2832. Maximal Range That Each Element Is Maximum In It](https://leetcode.com/problems/maximal-range-that-each-element-is-maximum-in-it/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func MaximalRangeThatEachElementIsMaximumInIt(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #2832: Maximal Range That Each Element Is Maximum in It
// https://leetcode.com/problems/maximal-range-that-each-element-is-maximum-in-it/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximalRangeThatEachElementIsMaximumInIt(nums []int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)

	// Previous greater element
  // Membuat slice untuk menyimpan hasil
	prev := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prev[i] = stack[len(stack)-1]
		} else {
			prev[i] = -1
		}
		stack = append(stack, i)
	}

	// Next greater element
  // Membuat slice untuk menyimpan hasil
	next := make([]int, n)
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			next[i] = stack[len(stack)-1]
		} else {
			next[i] = n
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		result[i] = next[i] - prev[i] - 1
	}

	return result
}

func main() {
	fmt.Println(MaximalRangeThatEachElementIsMaximumInIt([]int{1, 5, 4, 3, 6}))
	fmt.Println(MaximalRangeThatEachElementIsMaximumInIt([]int{1, 2, 1}))
}
```
