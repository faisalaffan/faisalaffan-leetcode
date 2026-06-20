# 1549 — The Most Recent Orders For Each Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func MostRecentOrders(products map[int]string, orders []struct{ orderID, productID int; date string }) []recentOrderInfo`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1549: The Most Recent Orders for Each Product
// https://leetcode.com/problems/the-most-recent-orders-for-each-product/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// SQL problem: for each product, find the most recent order(s).
	// Tables: Products(product_id, product_name), Orders(order_id, product_id, order_date)

	products := map[int]string{
		1: "Product A",
		2: "Product B",
		3: "Product C",
	}

	orders := []struct{ orderID, productID int; date string }{
		{1, 1, "2020-07-31"},
		{2, 1, "2020-07-30"},
		{3, 1, "2020-07-29"},
		{4, 2, "2020-07-31"},
		{5, 2, "2020-07-30"},
		{6, 3, "2020-07-31"},
	}

	result := MostRecentOrders(products, orders)
	fmt.Println("Most recent orders per product:")
	for _, r := range result {
		fmt.Printf("  %s: Order %d on %s\n", r.productName, r.orderID, r.date)
	}
}

type recentOrderInfo struct {
	productName string
	orderID     int
	date        string
}

func MostRecentOrders(products map[int]string, orders []struct{ orderID, productID int; date string }) []recentOrderInfo {
	// Group orders by product
  // HashMap: O(1) lookup
	productOrders := make(map[int][]struct{ orderID int; date string })
	for _, o := range orders {
		productOrders[o.productID] = append(productOrders[o.productID], struct{ orderID int; date string }{o.orderID, o.date})
	}

	// Find most recent order date per product
  // HashMap: O(1) lookup
	productRecent := make(map[int]string)
	for pid, ords := range productOrders {
  // Custom sort
		sort.Slice(ords, func(i, j int) bool {
			return ords[i].date > ords[j].date
		})
		productRecent[pid] = ords[0].date
	}

	// Collect orders that match the most recent date for their product
	result := make([]recentOrderInfo, 0)
	for pid, pname := range products {
		recentDate, ok := productRecent[pid]
		if !ok {
			continue
		}
		for _, o := range orders {
			if o.productID == pid && o.date == recentDate {
				result = append(result, recentOrderInfo{pname, o.orderID, o.date})
			}
		}
	}

	return result
}
```
