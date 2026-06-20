# 1128 — Number Of Equivalent Domino Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func numEquivDominoPairs(dominoes [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1128: Number of Equivalent Domino Pairs
// https://leetcode.com/problems/number-of-equivalent-domino-pairs/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}})) // 1
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}}))  // 3
}

// LeetCode submission: numEquivDominoPairs
func numEquivDominoPairs(dominoes [][]int) int {
  // HashMap: O(1) lookup
	count := make(map[[2]int]int)
	ans := 0
	for _, d := range dominoes {
		a, b := d[0], d[1]
		if a > b {
			a, b = b, a
		}
		key := [2]int{a, b}
		ans += count[key]
		count[key]++
	}
	return ans
}
```
