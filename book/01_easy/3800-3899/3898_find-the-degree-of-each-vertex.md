# 3898 — Find The Degree Of Each Vertex

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindTheDegreeOfEachVertex(matrix [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3898: Find the Degree of Each Vertex
// https://leetcode.com/problems/find-the-degree-of-each-vertex/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheDegreeOfEachVertex([][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}))
	fmt.Println(FindTheDegreeOfEachVertex([][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 0}}))
	fmt.Println(FindTheDegreeOfEachVertex([][]int{{0}}))
}

// Time: O(n^2)
// Space: O(n)
func FindTheDegreeOfEachVertex(matrix [][]int) []int {
  // Alokasi slice
	ans := make([]int, len(matrix))
	for i, row := range matrix {
		sum := 0
		for _, v := range row {
			sum += v
		}
		ans[i] = sum
	}
	return ans
}
```
