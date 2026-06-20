# 0496 — Next Greater Element I

## Deskripsi

**Soal:** [0496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n+m), Space: O(m)  
**Kompleksitas Ruang:** O(m)

**Algoritma:** Stack (tumpukan LIFO)

**Fungsi Solusi:** `func NextGreaterElementI(nums1, nums2 []int) []int`

## Solusi Go

```go
package main

// LeetCode #496: Next Greater Element I
// https://leetcode.com/problems/next-greater-element-i/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(m)
func NextGreaterElementI(nums1, nums2 []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	nextGreater := make(map[int]int)
	var stack []int
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			nextGreater[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(nums1))
	for i, v := range nums1 {
		if val, ok := nextGreater[v]; ok {
			result[i] = val
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	fmt.Println(NextGreaterElementI([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(NextGreaterElementI([]int{2, 4}, []int{1, 2, 3, 4}))
}
```
