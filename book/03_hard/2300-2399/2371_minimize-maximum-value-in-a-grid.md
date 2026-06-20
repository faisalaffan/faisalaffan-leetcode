# 2371 — Minimize Maximum Value In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minMaxValue(grid [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2371: Minimize Maximum Value in a Grid
// https://leetcode.com/problems/minimize-maximum-value-in-a-grid/
// Difficulty: Hard [Paid]
//
// Given a grid with positive integers, assign a new positive integer to each
// cell such that for each row and column, the assigned values maintain the
// same relational order as the original values (i.e., larger original value
// gets larger assigned value). Minimize the maximum assigned value.
//
// Approach: Sort cells by value. Process equal-value groups. For each cell,
// assigned value = max(rowMax[r], colMax[c]) + 1. After the group, update
// rowMax and colMax for the group.

import (
	"fmt"
	"sort"
)

func minMaxValue(grid [][]int) [][]int {
	rows := len(grid)
	if rows == 0 {
		return [][]int{}
	}
	cols := len(grid[0])

	type cell struct {
		val, r, c int
	}
	cells := make([]cell, 0, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cells = append(cells, cell{grid[i][j], i, j})
		}
	}

  // Custom sort
	sort.Slice(cells, func(i, j int) bool {
		return cells[i].val < cells[j].val
	})

  // Matriks 2D
	result := make([][]int, rows)
  // Range loop
	for i := range result {
		result[i] = make([]int, cols)
	}

  // Alokasi slice
	rowMax := make([]int, rows)
  // Alokasi slice
	colMax := make([]int, cols)

	i := 0
	for i < len(cells) {
		j := i
		for j < len(cells) && cells[j].val == cells[i].val {
			j++
		}

		// For each cell in the equal-value group, compute assigned value
		// but don't update rowMax/colMax until the entire group is done.
		type assign struct{ r, c, v int }
		assignments := make([]assign, 0, j-i)
		for k := i; k < j; k++ {
			r, c := cells[k].r, cells[k].c
			v := max(rowMax[r], colMax[c]) + 1
			result[r][c] = v
			assignments = append(assignments, assign{r, c, v})
		}

		// Now update rowMax and colMax for this group
		for _, a := range assignments {
			if a.v > rowMax[a.r] {
				rowMax[a.r] = a.v
			}
			if a.v > colMax[a.c] {
				colMax[a.c] = a.v
			}
		}

		i = j
	}

	return result
}

func main() {
	// Example 1
	grid1 := [][]int{
		{3, 1},
		{2, 5},
	}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid1) {
		fmt.Println(row)
	}

	// Example 2: equal values in same row
	grid2 := [][]int{
		{10, 10, 10},
	}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid2) {
		fmt.Println(row)
	}

	// Single cell
	grid3 := [][]int{{7}}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid3) {
		fmt.Println(row)
	}

	// 3x3 all distinct
	grid4 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid4) {
		fmt.Println(row)
	}

	// Multiple equal values
	grid5 := [][]int{
		{5, 5, 3},
		{5, 5, 4},
	}
	fmt.Println("Result:")
	for _, row := range minMaxValue(grid5) {
		fmt.Println(row)
	}
}
```
