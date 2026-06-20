# 2985 — Calculate Compressed Mean

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func CalculateCompressedMean(data [][]int) float64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2985: Calculate Compressed Mean
// https://leetcode.com/problems/calculate-compressed-mean/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent weighted average calculation.

import "fmt"

func main() {
	// LeetCode name: calculateCompressedMean
	// Input: item_count, order_occurrences pairs
	fmt.Println(CalculateCompressedMean([][]int{{1, 500}, {2, 1000}, {3, 800}, {4, 200}}))
	fmt.Println(CalculateCompressedMean([][]int{{1, 5}, {2, 10}, {3, 5}})) // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: calculateCompressedMean
// data[i] = [item_count, order_occurrences]
func CalculateCompressedMean(data [][]int) float64 {
	var totalItems, totalOrders float64
	for _, row := range data {
		totalItems += float64(row[0] * row[1])
		totalOrders += float64(row[1])
	}
	result := totalItems / totalOrders
	// Round to 2 decimal places
	return float64(int(result*100+0.5)) / 100
}
```
