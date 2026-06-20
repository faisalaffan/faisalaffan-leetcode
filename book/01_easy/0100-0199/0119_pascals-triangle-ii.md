# 0119 — Pascals Triangle Ii

## Deskripsi

**Soal:** [0119. Pascals Triangle Ii](https://leetcode.com/problems/pascals-triangle-ii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(rowIndex^2)  
**Kompleksitas Ruang:** O(rowIndex)

**Algoritma:** —

**Fungsi Solusi:** `func GetRow(rowIndex int) []int`

## Solusi Go

```go
package main

// LeetCode #119: Pascal's Triangle II
// https://leetcode.com/problems/pascals-triangle-ii/
// Difficulty: Easy

import "fmt"

// Time: O(rowIndex^2) | Space: O(rowIndex)
func GetRow(rowIndex int) []int {
  // Membuat slice untuk menyimpan hasil
	res := make([]int, rowIndex+1)
	res[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			res[j] += res[j-1]
		}
	}
	return res
}

func main() {
	fmt.Println(GetRow(3))
	fmt.Println(GetRow(0))
	fmt.Println(GetRow(4))
}
```
