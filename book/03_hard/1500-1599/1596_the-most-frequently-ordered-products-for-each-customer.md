# 1596 — The Most Frequently Ordered Products For Each Customer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mostFrequentProducts(orders []Order, products []Product) []resultRow1596
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1596: The Most Frequently Ordered Products for Each Customer
// https://leetcode.com/problems/the-most-frequently-ordered-products-for-each-customer/
// Difficulty: Hard
//
// For each customer, find the product(s) they ordered most frequently.
// If there are ties (multiple products ordered the same maximum number of times),
// include all of them.

// Order represents a customer order.
type Order struct {
	OrderID    int
	CustomerID int
	ProductID  int
}

// Product represents a product.
type Product struct {
	ProductID   int
	ProductName string
}

type resultRow1596 struct {
	CustomerID  int
	ProductName string
	ProductID   int
}

// mostFrequentProducts returns for each customer the product(s) they ordered most frequently.
func mostFrequentProducts(orders []Order, products []Product) []resultRow1596 {
	// Count orders per customer per product
	type cpKey struct {
		customerID int
		productID  int
	}
  // Membuat map (HashMap) — pencarian O(1)
	counts := make(map[cpKey]int)
	// Also track max per customer
  // Membuat map (HashMap) — pencarian O(1)
	maxPerCustomer := make(map[int]int)

	for _, o := range orders {
		key := cpKey{customerID: o.CustomerID, productID: o.ProductID}
		counts[key]++
		c := counts[key]
		if c > maxPerCustomer[o.CustomerID] {
			maxPerCustomer[o.CustomerID] = c
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	productName := make(map[int]string)
	for _, p := range products {
		productName[p.ProductID] = p.ProductName
	}

	var results []resultRow1596

	for key, cnt := range counts {
		if cnt == maxPerCustomer[key.customerID] {
			results = append(results, resultRow1596{
				CustomerID:  key.customerID,
				ProductName: productName[key.productID],
				ProductID:   key.productID,
			})
		}
	}

	// Sort results by customer ID ascending, product name ascending
  // Custom sort dengan comparator
	sort.Slice(results, func(i, j int) bool {
		if results[i].CustomerID != results[j].CustomerID {
			return results[i].CustomerID < results[j].CustomerID
		}
		return results[i].ProductName < results[j].ProductName
	})

	return results
}

func main() {
	orders := []Order{
		{1, 1, 10},
		{2, 1, 20},
		{3, 1, 10},
		{4, 2, 30},
		{5, 2, 30},
		{6, 2, 20},
		{7, 3, 10},
		{8, 3, 20},
		{9, 3, 30},
	}
	products := []Product{
		{10, "Widget"},
		{20, "Gadget"},
		{30, "Doohickey"},
	}

	results := mostFrequentProducts(orders, products)
	for _, r := range results {
		fmt.Printf("Customer %d | %s (ID %d)\n", r.CustomerID, r.ProductName, r.ProductID)
	}
	// Expected:
	// Customer 1 | Widget (ID 10)  (ordered 2x vs Gadget 1x)
	// Customer 2 | Doohickey (ID 30)  (ordered 2x vs Gadget 1x)
	// Customer 3 | Widget, Gadget, Doohickey (each ordered 1x)
}
```
