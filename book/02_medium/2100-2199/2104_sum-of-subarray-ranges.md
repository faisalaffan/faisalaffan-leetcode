# 2104 — Sum Of Subarray Ranges

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func subArrayRanges(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Stack, Monotonic Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2104: Sum of Subarray Ranges
// https://leetcode.com/problems/sum-of-subarray-ranges/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	var result int64 = 0

	// For each element, count subarrays where it's the max and where it's the min
	// Use monotonic stack to find prev/next greater/smaller

	// As max: prevGreater, nextGreater
  // Alokasi slice
	prevGreater := make([]int, n)
  // Alokasi slice
	nextGreater := make([]int, n)
  // Range loop
	for i := range prevGreater {
		prevGreater[i] = -1
		nextGreater[i] = n
	}

	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			nextGreater[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prevGreater[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// As min: prevSmaller, nextSmaller
  // Alokasi slice
	prevSmaller := make([]int, n)
  // Alokasi slice
	nextSmaller := make([]int, n)
  // Range loop
	for i := range prevSmaller {
		prevSmaller[i] = -1
		nextSmaller[i] = n
	}

	stack = []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			nextSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			prevSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		// Contribution as max
		leftMax := i - prevGreater[i]
		rightMax := nextGreater[i] - i
		result += int64(nums[i]) * int64(leftMax) * int64(rightMax)

		// Contribution as min
		leftMin := i - prevSmaller[i]
		rightMin := nextSmaller[i] - i
		result -= int64(nums[i]) * int64(leftMin) * int64(rightMin)
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", subArrayRanges([]int{1, 2, 3}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", subArrayRanges([]int{1, 3, 3}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", subArrayRanges([]int{4, -2, -3, 4, 1}))
	// Expected: 59
}
```
