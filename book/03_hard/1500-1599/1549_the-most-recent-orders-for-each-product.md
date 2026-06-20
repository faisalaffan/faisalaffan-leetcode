# 1549 — The Most Recent Orders For Each Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func mostRecentOrdersForEachProduct(products []Product, orders []Order) []resultRow1549`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1549: The Most Recent Orders for Each Product
// https://leetcode.com/problems/the-most-recent-orders-for-each-product/
// Difficulty: Hard
//
// For each product, find the most recent order(s) -- if multiple orders
// share the most recent date, include all of them.

// Order represents a customer order line.
type Order struct {
	OrderID   int
	OrderDate string
	ProductID int
	Quantity  int
}

// Product represents a product.
type Product struct {
	ProductID   int
	ProductName string
}

type resultRow1549 struct {
	ProductName string
	ProductID   int
	OrderID     int
	OrderDate   string
}

// mostRecentOrdersForEachProduct returns for each product its most recent order(s).
func mostRecentOrdersForEachProduct(products []Product, orders []Order) []resultRow1549 {
	// Group orders by product ID
  // HashMap: O(1) lookup
	ordersByProduct := make(map[int][]Order)
	for _, o := range orders {
		ordersByProduct[o.ProductID] = append(ordersByProduct[o.ProductID], o)
	}

  // HashMap: O(1) lookup
	productMap := make(map[int]string)
	for _, p := range products {
		productMap[p.ProductID] = p.ProductName
	}

	var results []resultRow1549

	for pid, ords := range ordersByProduct {
		// Find the most recent date
		maxDate := ""
		for _, o := range ords {
			if o.OrderDate > maxDate {
				maxDate = o.OrderDate
			}
		}

		// Collect all orders with that date
		for _, o := range ords {
			if o.OrderDate == maxDate {
				results = append(results, resultRow1549{
					ProductName: productMap[pid],
					ProductID:   pid,
					OrderID:     o.OrderID,
					OrderDate:   o.OrderDate,
				})
			}
		}
	}

	// Sort results by product name ascending, order ID ascending
  // Custom sort
	sort.Slice(results, func(i, j int) bool {
		if results[i].ProductName != results[j].ProductName {
			return results[i].ProductName < results[j].ProductName
		}
		return results[i].OrderID < results[j].OrderID
	})

	return results
}

func main() {
	products := []Product{
		{1, "Widget"},
		{2, "Gadget"},
	}
	orders := []Order{
		{101, "2023-01-01", 1, 5},
		{102, "2023-01-02", 1, 3},
		{103, "2023-01-02", 2, 2},
		{104, "2023-01-03", 2, 1},
	}

	results := mostRecentOrdersForEachProduct(products, orders)
	for _, r := range results {
		fmt.Printf("%s (ID %d) | Order %d | %s\n", r.ProductName, r.ProductID, r.OrderID, r.OrderDate)
	}
	// Expected:
	// Gadget (ID 2) | Order 104 | 2023-01-03
	// Widget (ID 1) | Order 102 | 2023-01-02
}
```
