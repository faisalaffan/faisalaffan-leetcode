# 1174 — Immediate Food Delivery Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func immediateFoodDeliveryII(deliveries []delivery) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1174: Immediate Food Delivery II
// https://leetcode.com/problems/immediate-food-delivery-ii/
// Difficulty: Medium

// For each customer's first order, find percentage that is "immediate"
// (order_date == customer_pref_delivery_date).

// Time: O(n)
// Space: O(n)

type delivery struct {
	ID         int
	customerID int
	orderDate  int
	prefDate   int
}

func immediateFoodDeliveryII(deliveries []delivery) float64 {
	// For each customer, track first order (earliest orderDate, tie-break by smallest ID)
	type firstOrder struct {
		date     int
		prefDate int
	}
  // HashMap: O(1) lookup
	custFirst := make(map[int]firstOrder)

	for _, d := range deliveries {
		fo, exists := custFirst[d.customerID]
		if !exists || d.orderDate < fo.date {
			custFirst[d.customerID] = firstOrder{d.orderDate, d.prefDate}
		}
	}

	immediate := 0
	total := 0
	for _, fo := range custFirst {
		total++
		if fo.date == fo.prefDate {
			immediate++
		}
	}

	if total == 0 {
		return 0
	}
	return math.Round(float64(immediate)/float64(total)*100) / 100
}

func main() {
	d1 := []delivery{
		{1, 1, 20190801, 20190801},
		{2, 2, 20190802, 20190803},
		{3, 1, 20190803, 20190804},
	}
	fmt.Printf("%.2f (expected: 0.50)\n", immediateFoodDeliveryII(d1))

	d2 := []delivery{
		{1, 1, 20190801, 20190801},
		{2, 1, 20190802, 20190802},
	}
	fmt.Printf("%.2f (expected: 1.00)\n", immediateFoodDeliveryII(d2))
}
```
