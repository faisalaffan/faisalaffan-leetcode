# 1710 — Maximum Units On A Truck

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaximumUnits(boxTypes [][]int, truckSize int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(log n) (for sorting)  |  **Ruang:** O(log n) (for sorting)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1710: Maximum Units on a Truck
// https://leetcode.com/problems/maximum-units-on-a-truck/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func MaximumUnits(boxTypes [][]int, truckSize int) int {
  // Custom sort
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	totalUnits := 0
	for _, box := range boxTypes {
		take := box[0]
		if truckSize < take {
			take = truckSize
		}
		totalUnits += take * box[1]
		truckSize -= take
		if truckSize == 0 {
			break
		}
	}
	return totalUnits
}

func main() {
	fmt.Println(MaximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println(MaximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
}
```
