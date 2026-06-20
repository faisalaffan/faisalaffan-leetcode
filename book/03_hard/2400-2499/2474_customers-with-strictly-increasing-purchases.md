# 2474 — Customers With Strictly Increasing Purchases

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func strictlyIncreasingPurchases(orders []Order) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Trie

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2474: Customers With Strictly Increasing Purchases
// https://leetcode.com/problems/customers-with-strictly-increasing-purchases/
// Difficulty: Hard [Paid] (SQL)
//
// Find customers who made purchases in at least two consecutive years
// with strictly increasing purchase amounts year over year.
//
// Approach: Simulate the SQL query logic in Go.

import (
	"fmt"
)

// Order represents a customer order
type Order struct {
	OrderID       int
	CustomerID    int
	OrderDate     string // "YYYY-MM-DD"
	Price         int
}

func main() {
	// Example
	orders := []Order{
		{1, 1, "2020-06-01", 100},
		{2, 1, "2021-07-01", 150},
		{3, 2, "2020-01-01", 200},
		{4, 2, "2021-02-01", 250},
		{5, 2, "2022-03-01", 300},
		{6, 3, "2020-05-01", 50},
		{7, 3, "2021-06-01", 60},
		{8, 3, "2022-07-01", 55},
		{9, 4, "2020-10-01", 500},
		{10, 4, "2021-11-01", 400},
	}

	fmt.Println(strictlyIncreasingPurchases(orders))
}

func strictlyIncreasingPurchases(orders []Order) []int {
	// Group min purchase per customer per year
	type yearAmount struct {
		year   int
		amount int
	}
  // HashMap: O(1) lookup
	customerYears := make(map[int][]yearAmount)

	for _, o := range orders {
		year := 0
		fmt.Sscanf(o.OrderDate[:4], "%d", &year)
		customerYears[o.CustomerID] = append(customerYears[o.CustomerID], yearAmount{year, o.Price})
	}

	var result []int
	for cid, entries := range customerYears {
		if len(entries) < 2 {
			continue
		}
		// Find min per year
  // HashMap: O(1) lookup
		minByYear := make(map[int]int)
		for _, e := range entries {
			if v, ok := minByYear[e.year]; !ok || e.amount < v {
				minByYear[e.year] = e.amount
			}
		}
		// Sort years
  // Alokasi slice
		years := make([]int, 0, len(minByYear))
		for y := range minByYear {
			years = append(years, y)
		}
		// Simple bubble sort for small sets
  // Linear scan O(n)
		for i := 0; i < len(years); i++ {
			for j := i + 1; j < len(years); j++ {
				if years[i] > years[j] {
					years[i], years[j] = years[j], years[i]
				}
			}
		}

		if len(years) < 2 {
			continue
		}
		// Check strictly increasing
		strict := true
		for i := 1; i < len(years); i++ {
			if years[i] != years[i-1]+1 || minByYear[years[i]] <= minByYear[years[i-1]] {
				strict = false
				break
			}
		}
		if strict {
			result = append(result, cid)
		}
	}

	// Sort result
  // Linear scan O(n)
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}
```
