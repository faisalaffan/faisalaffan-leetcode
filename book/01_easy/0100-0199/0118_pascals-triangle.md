# 0118 — Pascals Triangle

## Deskripsi

**Soal:** [0118. Pascals Triangle](https://leetcode.com/problems/pascals-triangle/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(numRows^2)  
**Kompleksitas Ruang:** O(numRows^2)

**Algoritma:** —

**Fungsi Solusi:** `func Generate(numRows int) [][]int`

## Solusi Go

```go
package main

// LeetCode #118: Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/
// Difficulty: Easy

import "fmt"

// Time: O(numRows^2) | Space: O(numRows^2)
func Generate(numRows int) [][]int {
  // Membuat slice 2D untuk DP/tabel
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

func main() {
	fmt.Println(Generate(5))
	fmt.Println(Generate(1))
}
```
