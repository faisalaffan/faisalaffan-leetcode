# 2893 — Calculate Orders Within Each Interval

## Deskripsi

**Soal:** [2893. Calculate Orders Within Each Interval](https://leetcode.com/problems/calculate-orders-within-each-interval/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CalculateOrdersWithinEachInterval(orders []Order, interval int) []int`

## Solusi Go

```go
package main

// LeetCode #2893: Calculate Orders Within Each Interval
// https://leetcode.com/problems/calculate-orders-within-each-interval/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

type Order struct {
	Time   int
	Amount int
}

func CalculateOrdersWithinEachInterval(orders []Order, interval int) []int {
	if len(orders) == 0 {
		return []int{}
	}

	// Group orders by interval
	maxTime := orders[len(orders)-1].Time
	bucketCount := maxTime/interval + 1
  // Membuat slice untuk menyimpan hasil
	buckets := make([]int, bucketCount)

	for _, o := range orders {
		idx := o.Time / interval
		buckets[idx] += o.Amount
	}

	return buckets
}

func main() {
	orders := []Order{
		{0, 10}, {1, 20}, {4, 30}, {6, 40},
	}
	fmt.Println(CalculateOrdersWithinEachInterval(orders, 3))

	orders2 := []Order{
		{0, 5}, {2, 10},
	}
	fmt.Println(CalculateOrdersWithinEachInterval(orders2, 5))
}
```
