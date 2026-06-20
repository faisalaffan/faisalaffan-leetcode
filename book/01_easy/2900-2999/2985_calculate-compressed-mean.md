# 2985 — Calculate Compressed Mean

## Deskripsi

**Soal:** [2985. Calculate Compressed Mean](https://leetcode.com/problems/calculate-compressed-mean/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

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
