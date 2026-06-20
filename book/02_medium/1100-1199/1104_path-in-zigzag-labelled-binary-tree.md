# 1104 — Path In Zigzag Labelled Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func pathInZigZagTree(label int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(log n)


## 💻 Solusi Go

```go
package main

// LeetCode #1104: Path In Zigzag Labelled Binary Tree
// https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/
// Difficulty: Medium
//
// Approach: Find level, compute reverse label, traverse to root
// Time: O(log n)
// Space: O(log n)

import "fmt"

func main() {
	fmt.Println(pathInZigZagTree(14)) // [1,3,4,14]
	fmt.Println(pathInZigZagTree(26)) // [1,2,6,10,26]
}

func pathInZigZagTree(label int) []int {
  // Alokasi slice
	result := make([]int, 0)

	for label > 0 {
		result = append(result, label)
		level := 0
		for (1 << level) <= label {
			level++
		}
		level--

		// In zigzag levels, the "position" is reversed
		// Min and max of this level
		minVal := 1 << level
		maxVal := (1 << (level + 1)) - 1
		// The parent of label (in zigzag) needs to find the "mirror" position
		parent := minVal + maxVal - label
		label = parent / 2
	}

	// Reverse to get root-to-leaf
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
```
