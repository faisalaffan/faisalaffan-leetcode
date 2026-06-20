# 1867 — Orders With Maximum Quantity Above Average

## Deskripsi

**Soal:** [1867. Orders With Maximum Quantity Above Average](https://leetcode.com/problems/orders-with-maximum-quantity-above-average/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1867: Orders With Maximum Quantity Above Average
// https://leetcode.com/problems/orders-with-maximum-quantity-above-average/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// Sample: order_id, quantity
	orders := [][]int{{1, 10}, {2, 5}, {3, 8}, {4, 3}, {5, 12}}
	fmt.Println(OrdersAboveAverage(orders))
}

// Time: O(n), Space: O(1)
func OrdersAboveAverage(orders [][]int) int {
	if len(orders) == 0 {
		return 0
	}
	sum := 0
	for _, o := range orders {
		sum += o[1]
	}
	avg := float64(sum) / float64(len(orders))
	count := 0
	for _, o := range orders {
		if float64(o[1]) > avg {
			count++
		}
	}
	// Find max quantity among orders above average
	maxQty := 0
	for _, o := range orders {
		if float64(o[1]) > avg && o[1] > maxQty {
			maxQty = o[1]
		}
	}
	return maxQty
}
```
