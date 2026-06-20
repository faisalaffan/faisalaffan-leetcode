# 0496 — Next Greater Element I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func NextGreaterElementI(nums1, nums2 []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Stack

**Waktu:** O(n+m), Space: O(m)  |  **Ruang:** O(m)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #496: Next Greater Element I
// https://leetcode.com/problems/next-greater-element-i/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(m)
func NextGreaterElementI(nums1, nums2 []int) []int {
  // HashMap: O(1) lookup
	nextGreater := make(map[int]int)
	var stack []int
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			nextGreater[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
  // Alokasi slice
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
