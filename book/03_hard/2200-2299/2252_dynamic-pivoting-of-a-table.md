# 2252 — Dynamic Pivoting Of A Table

## Deskripsi

**Soal:** [2252. Dynamic Pivoting Of A Table](https://leetcode.com/problems/dynamic-pivoting-of-a-table/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func NewPivotTable() *PivotTable`

## Solusi Go

```go
package main

// LeetCode #2252: Dynamic Pivoting of a Table
// https://leetcode.com/problems/dynamic-pivoting-of-a-table/
// Difficulty: Hard [Paid]
//
// Given a table with columns (product_id, store, price),
// pivot it so that each store becomes a column and each row is a product_id.
// The store names are dynamic (not known in advance).

import (
	"fmt"
)

// PivotTable represents a pivoted in-memory table.
type PivotTable struct {
	// map[product_id][store] = price
	Data map[int]map[string]int
}

// NewPivotTable creates a new pivot table.
func NewPivotTable() *PivotTable {
	return &PivotTable{Data: make(map[int]map[string]int)}
}

// AddRow adds a (product_id, store, price) record.
func (pt *PivotTable) AddRow(productID int, store string, price int) {
	if pt.Data[productID] == nil {
		pt.Data[productID] = make(map[string]int)
	}
	pt.Data[productID][store] = price
}

// Pivot returns the pivoted representation:
// rows = product_ids, columns = stores (sorted), values = prices.
func (pt *PivotTable) Pivot() ([]string, map[int]map[string]int) {
	// collect all stores
  // Membuat map untuk pencarian O(1): key → value
	storeSet := make(map[string]bool)
	for _, stores := range pt.Data {
		for s := range stores {
			storeSet[s] = true
		}
	}

  // Membuat slice untuk menyimpan hasil
	stores := make([]string, 0, len(storeSet))
	for s := range storeSet {
		stores = append(stores, s)
	}
	// sort stores alphabetically (simple bubble for small set)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(stores); i++ {
		for j := i + 1; j < len(stores); j++ {
			if stores[j] < stores[i] {
				stores[i], stores[j] = stores[j], stores[i]
			}
		}
	}

	return stores, pt.Data
}

func main() {
	pt := NewPivotTable()
	pt.AddRow(1, "StoreA", 10)
	pt.AddRow(1, "StoreB", 20)
	pt.AddRow(2, "StoreA", 15)
	pt.AddRow(2, "StoreC", 25)
	pt.AddRow(3, "StoreB", 30)

	stores, data := pt.Pivot()
	fmt.Println("Stores:", stores)
	for pid := 1; pid <= 3; pid++ {
		fmt.Printf("Product %d: %v\n", pid, data[pid])
	}
}
```
