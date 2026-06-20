# 3477 — Fruits Into Baskets Ii

## Deskripsi

**Soal:** [3477. Fruits Into Baskets Ii](https://leetcode.com/problems/fruits-into-baskets-ii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3477: Fruits Into Baskets II
// https://leetcode.com/problems/fruits-into-baskets-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FruitsIntoBasketsIi([]int{4, 2, 5}, []int{3, 5, 4}))
	fmt.Println(FruitsIntoBasketsIi([]int{3, 6, 1}, []int{6, 4, 7}))
}

// FruitsIntoBasketsIi counts fruits that cannot be placed into baskets.
// Each fruit i can go into basket j if fruits[i] <= baskets[j].
// Time: O(n * m). Space: O(1).
func FruitsIntoBasketsIi(fruits []int, baskets []int) int {
  // Membuat slice untuk menyimpan hasil
	used := make([]bool, len(baskets))
	unplaced := 0
	for _, f := range fruits {
		placed := false
		for j, b := range baskets {
			if !used[j] && f <= b {
				used[j] = true
				placed = true
				break
			}
		}
		if !placed {
			unplaced++
		}
	}
	return unplaced
}
```
