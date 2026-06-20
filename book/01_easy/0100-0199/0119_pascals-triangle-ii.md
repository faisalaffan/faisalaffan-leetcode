# 0119 — Pascals Triangle Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func GetRow(rowIndex int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(rowIndex^2)  |  **Ruang:** O(rowIndex)


## 💻 Solusi Go

```go
package main

// LeetCode #119: Pascal's Triangle II
// https://leetcode.com/problems/pascals-triangle-ii/
// Difficulty: Easy

import "fmt"

// Time: O(rowIndex^2) | Space: O(rowIndex)
func GetRow(rowIndex int) []int {
  // Alokasi slice
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
