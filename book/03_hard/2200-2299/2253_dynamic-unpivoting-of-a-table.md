# 2253 — Dynamic Unpivoting Of A Table

## Deskripsi

**Soal:** [2253. Dynamic Unpivoting Of A Table](https://leetcode.com/problems/dynamic-unpivoting-of-a-table/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func NewUnpivotTable() *UnpivotTable`

## Solusi Go

```go
package main

// LeetCode #2253: Dynamic Unpivoting of a Table
// https://leetcode.com/problems/dynamic-unpivoting-of-a-table/
// Difficulty: Hard [Paid]
//
// Given a pivoted table where columns are dynamic store names,
// unpivot it back to (product_id, store, price) format.

import (
	"fmt"
	"sort"
)

// UnpivotTable handles dynamic unpivoting.
type UnpivotTable struct {
	// map[product_id][store] = price
	Data map[int]map[string]int
}

// NewUnpivotTable creates a new unpivot table.
func NewUnpivotTable() *UnpivotTable {
	return &UnpivotTable{Data: make(map[int]map[string]int)}
}

// LoadPivoted adds a row from pivoted format: productID with store prices.
func (ut *UnpivotTable) LoadPivoted(productID int, prices map[string]int) {
	for store, price := range prices {
		if ut.Data[productID] == nil {
			ut.Data[productID] = make(map[string]int)
		}
		ut.Data[productID][store] = price
	}
}

// Unpivot converts pivoted data to rows of (product_id, store, price).
func (ut *UnpivotTable) Unpivot() [][3]interface{} {
	var rows [][3]interface{}

	for productID, stores := range ut.Data {
		// sort stores for deterministic output
  // Membuat slice untuk menyimpan hasil
		storeNames := make([]string, 0, len(stores))
		for s := range stores {
			storeNames = append(storeNames, s)
		}
		sort.Strings(storeNames)

		for _, store := range storeNames {
			rows = append(rows, [3]interface{}{productID, store, stores[store]})
		}
	}

	return rows
}

func main() {
	ut := NewUnpivotTable()
	ut.LoadPivoted(1, map[string]int{"StoreA": 10, "StoreB": 20})
	ut.LoadPivoted(2, map[string]int{"StoreA": 15, "StoreC": 25})
	ut.LoadPivoted(3, map[string]int{"StoreB": 30})

	rows := ut.Unpivot()
	fmt.Println("Unpivoted rows:")
	for _, row := range rows {
		fmt.Printf("  Product %d, Store %s, Price %d\n", row[0].(int), row[1].(string), row[2].(int))
	}
}
```
