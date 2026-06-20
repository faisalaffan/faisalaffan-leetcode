# 1398 — Customers Who Bought Products A And B But Not C

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func customersABnotC(customers []struct {
	customerID   int
	customerName string
}, orders []struct {
	orderID     int
	customerID  int
	productName string
}) []customerResult
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n) where n = number of orders  
**Kompleksitas Ruang:** O(k) where k = number of customers

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1398: Customers Who Bought Products A and B but Not C
// https://leetcode.com/problems/customers-who-bought-products-a-and-b-but-not-c/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// SQL problem - simulating in Go
	result := customersABnotC(
		[]struct {
			customerID   int
			customerName string
		}{
			{1, "Daniel"},
			{2, "Diana"},
			{3, "Elizabeth"},
			{4, "John"},
		},
		[]struct {
			orderID     int
			customerID  int
			productName string
		}{
			{1, 1, "A"},
			{2, 1, "B"},
			{3, 1, "C"},
			{4, 2, "A"},
			{5, 2, "B"},
			{6, 3, "A"},
		},
	)
	for _, r := range result {
		fmt.Printf("%d %s\n", r.id, r.name)
	}
	// Should output: 2 Diana (bought A and B but not C)
	// Note: 1 Daniel bought A and B but also C, excluded
	// Note: 3 Elizabeth bought A but not B, excluded
}

type customerResult struct {
	id   int
	name string
}

// Time: O(n) where n = number of orders
// Space: O(k) where k = number of customers
func customersABnotC(customers []struct {
	customerID   int
	customerName string
}, orders []struct {
	orderID     int
	customerID  int
	productName string
}) []customerResult {
  // Membuat map (HashMap) — pencarian O(1)
	bought := make(map[int]map[string]bool)
  // Membuat map (HashMap) — pencarian O(1)
	customerNames := make(map[int]string)

	for _, c := range customers {
		customerNames[c.customerID] = c.customerName
	}

	for _, o := range orders {
		if bought[o.customerID] == nil {
			bought[o.customerID] = make(map[string]bool)
		}
		bought[o.customerID][o.productName] = true
	}

	var result []customerResult
	for _, c := range customers {
		products := bought[c.customerID]
		if products["A"] && products["B"] && !products["C"] {
			result = append(result, customerResult{c.customerID, c.customerName})
		}
	}

  // Custom sort dengan comparator
	sort.Slice(result, func(i, j int) bool {
		return result[i].id < result[j].id
	})

	return result
}
```
