# 1532 — The Most Recent Three Orders

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func RecentThreeOrders(orders []struct{ id, customerID int; date string; cost float64 }) []orderRec
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
	customerOrders := make(map[int][]struct{ id int; date string; cost float64 })
	for _, o := range orders {
		customerOrders[o.customerID] = append(customerOrders[o.customerID], struct{ id int; date string; cost float64 }{o.id, o.date, o.cost})
	}

	result := make([]orderRec, 0)
	for cid, ords := range customerOrders {
		// Sort by date descending
  // Custom sort dengan comparator
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
