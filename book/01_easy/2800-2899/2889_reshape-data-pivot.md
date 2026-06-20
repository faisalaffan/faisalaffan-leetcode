# 2889 — Reshape Data Pivot

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func ReshapeDataPivot(data [][]string) map[string]map[string]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2889: Reshape Data: Pivot
// https://leetcode.com/problems/reshape-data-pivot/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we pivot data from [month, city, temperature] to [month, city1, city2, ...].

import "fmt"

func main() {
	// LeetCode name: pivotTable
	// Input: [month, city, temperature]
	data := [][]string{
		{"January", "London", "5"},
		{"January", "Paris", "7"},
		{"February", "London", "6"},
		{"February", "Paris", "8"},
	}
	fmt.Println(ReshapeDataPivot(data))
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: pivotTable
func ReshapeDataPivot(data [][]string) map[string]map[string]string {
  // HashMap: O(1) lookup
	result := make(map[string]map[string]string)
	for _, row := range data {
		month, city, temp := row[0], row[1], row[2]
		if result[month] == nil {
			result[month] = make(map[string]string)
		}
		result[month][city] = temp
	}
	return result
}
```
