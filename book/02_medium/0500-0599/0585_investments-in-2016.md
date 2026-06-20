# 0585 — Investments In 2016

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindInvestmentSum(insurance [][]interface{}) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #585: Investments in 2016
// https://leetcode.com/problems/investments-in-2016/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Insurance records: {pid, tiv_2015, tiv_2016, lat, lon}
	insurance := [][]interface{}{
		{1, 100, 200, 10, 20},
		{2, 100, 300, 30, 40},
		{3, 200, 400, 50, 60},
		{4, 100, 500, 10, 20}, // same lat/lon as pid=1
	}
	fmt.Println(FindInvestmentSum(insurance))
}

func FindInvestmentSum(insurance [][]interface{}) float64 {
  // HashMap: O(1) lookup
	tiv2015 := make(map[float64]int)
  // HashMap: O(1) lookup
	locationCount := make(map[string]int)
  // HashMap: O(1) lookup
	tiv2016Sum := make(map[int]float64)

	for _, record := range insurance {
		pid := record[0].(int)
		tiv15 := record[1].(float64)
		tiv16 := record[2].(float64)
		lat := record[3].(float64)
		lon := record[4].(float64)

		tiv2015[tiv15]++
		locKey := fmt.Sprintf("%f,%f", lat, lon)
		locationCount[locKey]++
		tiv2016Sum[pid] = tiv16
	}

	total := 0.0
	for _, record := range insurance {
		pid := record[0].(int)
		tiv15 := record[1].(float64)
		lat := record[3].(float64)
		lon := record[4].(float64)
		locKey := fmt.Sprintf("%f,%f", lat, lon)

		if tiv2015[tiv15] > 1 && locationCount[locKey] == 1 {
			total += tiv2016Sum[pid]
		}
	}

	return total
}
```
