# 2362 — Generate The Invoice

## Deskripsi

**Soal:** [2362. Generate The Invoice](https://leetcode.com/problems/generate-the-invoice/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func GenerateTheInvoice(products []Product, purchases []Purchase) []Invoice`

## Solusi Go

```go
package main

// LeetCode #2362: Generate the Invoice
// https://leetcode.com/problems/generate-the-invoice/
// Difficulty: Hard [Paid]
//
// Given two tables: Products(product_id, price) and Purchases(invoice_id, product_id, quantity),
// return the invoice(s) with the maximum total price. If multiple invoices have the same max
// total, return all of them ordered by invoice_id.

import (
	"fmt"
	"sort"
)

type Product struct {
	ID    int
	Price int
}

type Purchase struct {
	InvoiceID int
	ProductID int
	Quantity  int
}

type Invoice struct {
	InvoiceID int
	Total     int
}

// GenerateTheInvoice returns invoices with the maximum total price.
func GenerateTheInvoice(products []Product, purchases []Purchase) []Invoice {
  // Membuat map untuk pencarian O(1): key → value
	priceMap := make(map[int]int)
	for _, p := range products {
		priceMap[p.ID] = p.Price
	}

  // Membuat map untuk pencarian O(1): key → value
	totals := make(map[int]int)
	for _, p := range purchases {
		totals[p.InvoiceID] += priceMap[p.ProductID] * p.Quantity
	}

	maxTotal := 0
	for _, v := range totals {
		if v > maxTotal {
			maxTotal = v
		}
	}

	var result []Invoice
	for id, total := range totals {
		if total == maxTotal {
			result = append(result, Invoice{InvoiceID: id, Total: total})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].InvoiceID < result[j].InvoiceID
	})

	return result
}

func main() {
	// Example 1
	products := []Product{
		{ID: 1, Price: 100},
		{ID: 2, Price: 200},
	}
	purchases := []Purchase{
		{InvoiceID: 1, ProductID: 1, Quantity: 2},
		{InvoiceID: 1, ProductID: 2, Quantity: 1},
		{InvoiceID: 2, ProductID: 1, Quantity: 1},
	}
	fmt.Println(GenerateTheInvoice(products, purchases))

	// Example 2: tie
	products2 := []Product{
		{ID: 1, Price: 50},
		{ID: 2, Price: 30},
	}
	purchases2 := []Purchase{
		{InvoiceID: 10, ProductID: 1, Quantity: 1},
		{InvoiceID: 20, ProductID: 2, Quantity: 1},
	}
	fmt.Println(GenerateTheInvoice(products2, purchases2))

	// Single invoice
	products3 := []Product{{ID: 5, Price: 75}}
	purchases3 := []Purchase{{InvoiceID: 7, ProductID: 5, Quantity: 4}}
	fmt.Println(GenerateTheInvoice(products3, purchases3))
}
```
