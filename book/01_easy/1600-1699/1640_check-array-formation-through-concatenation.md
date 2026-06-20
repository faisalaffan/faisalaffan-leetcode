# 1640 — Check Array Formation Through Concatenation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CanFormArray(arr []int, pieces [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1640: Check Array Formation Through Concatenation
// https://leetcode.com/problems/check-array-formation-through-concatenation/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func CanFormArray(arr []int, pieces [][]int) bool {
  // HashMap: O(1) lookup
	pos := make(map[int]int)
	for i, num := range arr {
		pos[num] = i
	}
	for _, piece := range pieces {
		first := piece[0]
		idx, ok := pos[first]
		if !ok {
			return false
		}
		for j := 1; j < len(piece); j++ {
			if idx+j >= len(arr) || arr[idx+j] != piece[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(CanFormArray([]int{15, 88}, [][]int{{88}, {15}}))
	fmt.Println(CanFormArray([]int{49, 18, 16}, [][]int{{16, 18, 49}}))
	fmt.Println(CanFormArray([]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}))
}
```
