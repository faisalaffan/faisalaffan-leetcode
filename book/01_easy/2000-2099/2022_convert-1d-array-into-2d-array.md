# 2022 — Convert 1D Array Into 2D Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ConvertOneDArrayIntoTwoDArray(original []int, m int, n int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n), Space: O(m*n)  |  **Ruang:** O(m*n)


## 💻 Solusi Go

```go
package main

// LeetCode #2022: Convert 1D Array Into 2D Array
// https://leetcode.com/problems/convert-1d-array-into-2d-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2, 3, 4}, 2, 2)) // [[1 2] [3 4]]
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2, 3}, 1, 3))    // [[1 2 3]]
	fmt.Println(ConvertOneDArrayIntoTwoDArray([]int{1, 2}, 1, 1))       // []
}

// Time: O(m*n), Space: O(m*n)
func ConvertOneDArrayIntoTwoDArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

  // Matriks 2D
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = original[i*n+j]
		}
	}
	return result
}
```
