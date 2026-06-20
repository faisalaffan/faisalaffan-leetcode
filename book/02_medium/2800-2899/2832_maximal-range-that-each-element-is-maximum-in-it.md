# 2832 — Maximal Range That Each Element Is Maximum In It

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximalRangeThatEachElementIsMaximumInIt(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2832: Maximal Range That Each Element Is Maximum in It
// https://leetcode.com/problems/maximal-range-that-each-element-is-maximum-in-it/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximalRangeThatEachElementIsMaximumInIt(nums []int) []int {
	n := len(nums)
  // Alokasi slice
	result := make([]int, n)

	// Previous greater element
  // Alokasi slice
	prev := make([]int, n)
  // Alokasi slice
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
  // Alokasi slice
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
