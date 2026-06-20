# 1532 — The Most Recent Three Orders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func RecentThreeOrders(orders []struct{ id, customerID int; date string; cost float64 }) []orderRec`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1532: The Most Recent Three Orders
// https://leetcode.com/problems/the-most-recent-three-orders/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// SQL problem: For each customer, find their 3 most recent orders.
	// Translated to Go.
	// Tables: Customers(customer_id, name), Orders(order_id, order_date, customer_id, cost)

	orders := []struct{ id, customerID int; date string; cost float64 }{
		{1, 1, "2020-07-31", 30.0},
		{2, 1, "2020-07-30", 40.0},
		{3, 1, "2020-07-29", 20.0},
		{4, 1, "2020-07-28", 50.0},
		{5, 2, "2020-07-31", 10.0},
		{6, 2, "2020-07-30", 15.0},
		{7, 3, "2020-07-31", 25.0},
	}

	result := RecentThreeOrders(orders)
	fmt.Println("Recent 3 orders per customer:")
	for _, r := range result {
		fmt.Printf("  Customer %d: Order %d on %s ($%.2f)\n", r.customerID, r.orderID, r.date, r.cost)
	}
}

type orderRec struct {
	customerID int
	orderID    int
	date       string
	cost       float64
}

func RecentThreeOrders(orders []struct{ id, customerID int; date string; cost float64 }) []orderRec {
	// Group orders by customer
  // HashMap: O(1) lookup
	customerOrders := make(map[int][]struct{ id int; date string; cost float64 })
	for _, o := range orders {
		customerOrders[o.customerID] = append(customerOrders[o.customerID], struct{ id int; date string; cost float64 }{o.id, o.date, o.cost})
	}

	result := make([]orderRec, 0)
	for cid, ords := range customerOrders {
		// Sort by date descending
  // Custom sort
		sort.Slice(ords, func(i, j int) bool {
			return ords[i].date > ords[j].date
		})
		// Take top 3
		for i := 0; i < 3 && i < len(ords); i++ {
			result = append(result, orderRec{cid, ords[i].id, ords[i].date, ords[i].cost})
		}
	}

	return result
}
```
