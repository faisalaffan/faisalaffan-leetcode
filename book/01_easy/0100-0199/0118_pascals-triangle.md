# 0118 — Pascals Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func Generate(numRows int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(numRows^2)  |  **Ruang:** O(numRows^2)


## 💻 Solusi Go

```go
package main

// LeetCode #118: Pascal's Triangle
// https://leetcode.com/problems/pascals-triangle/
// Difficulty: Easy

import "fmt"

// Time: O(numRows^2) | Space: O(numRows^2)
func Generate(numRows int) [][]int {
  // Matriks 2D
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
