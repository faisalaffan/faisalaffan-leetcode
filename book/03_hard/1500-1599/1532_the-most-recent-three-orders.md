# 1532 — The Most Recent Three Orders

## Deskripsi

**Soal:** [1532. The Most Recent Three Orders](https://leetcode.com/problems/the-most-recent-three-orders/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func mostRecentThreeOrders(customers []Customer, orders []Order) []resultRow1532`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1532: The Most Recent Three Orders
// https://leetcode.com/problems/the-most-recent-three-orders/
// Difficulty: Hard
//
// For each customer, find the most recent 3 orders.
// Implementation simulates the SQL query using Go data structures.

// Order represents a customer order.
type Order struct {
	OrderID    int
	OrderDate  string
	CustomerID int
}

// Customer represents a customer.
type Customer struct {
	CustomerID int
	Name       string
}

type resultRow1532 struct {
	CustomerName string
	OrderID      int
	OrderDate    string
}

// mostRecentThreeOrders returns for each customer their 3 most recent orders,
// ordered by customer name ascending, then order date descending, then order ID descending.
func mostRecentThreeOrders(customers []Customer, orders []Order) []resultRow1532 {
	// Group orders by customer ID
  // Membuat map untuk pencarian O(1): key → value
	ordersByCustomer := make(map[int][]Order)
	for _, o := range orders {
		ordersByCustomer[o.CustomerID] = append(ordersByCustomer[o.CustomerID], o)
	}

	// For each customer, sort orders by date descending, then order ID descending
  // Membuat map untuk pencarian O(1): key → value
	customerMap := make(map[int]string)
	for _, c := range customers {
		customerMap[c.CustomerID] = c.Name
	}

	var results []resultRow1532

	for cid, ords := range ordersByCustomer {
		sort.Slice(ords, func(i, j int) bool {
			if ords[i].OrderDate != ords[j].OrderDate {
				return ords[i].OrderDate > ords[j].OrderDate
			}
			return ords[i].OrderID > ords[j].OrderID
		})

		// Take top 3
		limit := 3
		if len(ords) < limit {
			limit = len(ords)
		}

		for _, o := range ords[:limit] {
			results = append(results, resultRow1532{
				CustomerName: customerMap[cid],
				OrderID:      o.OrderID,
				OrderDate:    o.OrderDate,
			})
		}
	}

	// Sort results by customer name ascending, order date descending, order ID descending
	sort.Slice(results, func(i, j int) bool {
		if results[i].CustomerName != results[j].CustomerName {
			return results[i].CustomerName < results[j].CustomerName
		}
		if results[i].OrderDate != results[j].OrderDate {
			return results[i].OrderDate > results[j].OrderDate
		}
		return results[i].OrderID > results[j].OrderID
	})

	return results
}

func main() {
	customers := []Customer{
		{1, "Alice"},
		{2, "Bob"},
	}
	orders := []Order{
		{101, "2023-01-01", 1},
		{102, "2023-01-02", 1},
		{103, "2023-01-03", 1},
		{104, "2023-01-04", 1},
		{201, "2023-01-01", 2},
		{202, "2023-01-02", 2},
	}

	results := mostRecentThreeOrders(customers, orders)
	for _, r := range results {
		fmt.Printf("%s | %d | %s\n", r.CustomerName, r.OrderID, r.OrderDate)
	}
	// Expected:
	// Alice | 104 | 2023-01-04
	// Alice | 103 | 2023-01-03
	// Alice | 102 | 2023-01-02
	// Bob   | 202 | 2023-01-02
	// Bob   | 201 | 2023-01-01
}
```
