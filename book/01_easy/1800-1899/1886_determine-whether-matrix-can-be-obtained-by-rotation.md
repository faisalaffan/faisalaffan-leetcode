# 1886 — Determine Whether Matrix Can Be Obtained By Rotation

## Deskripsi

**Soal:** [1886. Determine Whether Matrix Can Be Obtained By Rotation](https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func FindRotation(mat [][]int, target [][]int) bool`

## Solusi Go

```go
package main

// LeetCode #1886: Determine Whether Matrix Can Be Obtained by Rotation
// https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/
// Difficulty: Easy

import "fmt"

// Time: O(n^2), Space: O(1)
func FindRotation(mat [][]int, target [][]int) bool {
	for rotation := 0; rotation < 4; rotation++ {
		if equal(mat, target) {
			return true
		}
		mat = rotate(mat)
	}
	return false
}

func rotate(mat [][]int) [][]int {
	n := len(mat)
  // Membuat slice 2D untuk DP/tabel
	rotated := make([][]int, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rotated[i][j] = mat[n-1-j][i]
		}
	}
	return rotated
}

func equal(a, b [][]int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(FindRotation([][]int{{0, 1}, {1, 0}}, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(FindRotation([][]int{{0, 1}, {1, 1}}, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(FindRotation([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}))
}
```
