# 1428 — Leftmost Column With At Least A One

## Deskripsi

**Soal:** [1428. Leftmost Column With At Least A One](https://leetcode.com/problems/leftmost-column-with-at-least-a-one/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m + n) where m = rows, n = cols  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1428: Leftmost Column with at Least a One
// https://leetcode.com/problems/leftmost-column-with-at-least-a-one/
// Difficulty: Medium

import "fmt"

type BinaryMatrix struct {
	grid [][]int
}

func (bm BinaryMatrix) Get(row, col int) int {
	return bm.grid[row][col]
}

func (bm BinaryMatrix) Dimensions() []int {
	if len(bm.grid) == 0 {
		return []int{0, 0}
	}
	return []int{len(bm.grid), len(bm.grid[0])}
}

func main() {
	// Test case 1
	bm := BinaryMatrix{[][]int{{0, 0}, {1, 1}}}
	fmt.Println(leftMostColumnWithOne(bm)) // 0

	// Test case 2
	bm2 := BinaryMatrix{[][]int{{0, 0}, {0, 1}}}
	fmt.Println(leftMostColumnWithOne(bm2)) // 1

	// Test case 3
	bm3 := BinaryMatrix{[][]int{{0, 0}, {0, 0}}}
	fmt.Println(leftMostColumnWithOne(bm3)) // -1
}

// Time: O(m + n) where m = rows, n = cols
// Space: O(1)
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dim := binaryMatrix.Dimensions()
	m, n := dim[0], dim[1]

	row, col := 0, n-1
	leftmost := -1

	for row < m && col >= 0 {
		if binaryMatrix.Get(row, col) == 1 {
			leftmost = col
			col--
		} else {
			row++
		}
	}

	return leftmost
}
```
