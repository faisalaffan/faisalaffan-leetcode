# 1130 — Minimum Cost Tree From Leaf Values

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func mctFromLeafValues(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack, Monotonic Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1130: Minimum Cost Tree From Leaf Values
// https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/
// Difficulty: Medium
//
// Approach: Monotonic decreasing stack (greedy)
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(mctFromLeafValues([]int{6, 2, 4}))    // 32
	fmt.Println(mctFromLeafValues([]int{4, 11}))      // 44
}

func mctFromLeafValues(arr []int) int {
  // Alokasi slice
	stack := make([]int, 0)
	result := 0

	for _, v := range arr {
		for len(stack) > 0 && stack[len(stack)-1] <= v {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				result += mid * v
			} else {
				if stack[len(stack)-1] < v {
					result += mid * stack[len(stack)-1]
				} else {
					result += mid * v
				}
			}
		}
		stack = append(stack, v)
	}

	for len(stack) > 1 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result += last * stack[len(stack)-1]
	}

	return result
}
```
