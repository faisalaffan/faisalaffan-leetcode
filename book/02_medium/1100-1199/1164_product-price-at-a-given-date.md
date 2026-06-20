# 1164 — Product Price At A Given Date

## Deskripsi

**Soal:** [1164. Product Price At A Given Date](https://leetcode.com/problems/product-price-at-a-given-date/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func productPriceAtDate(products [][]int) []productPrice`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1164: Product Price at a Given Date
// https://leetcode.com/problems/product-price-at-a-given-date/
// Difficulty: Medium

// For each product, find its price on 2019-08-16.
// If no change after date, price is 10 (default).
// Find the most recent price change on or before the date.

// Time: O(n log n)
// Space: O(n)

type productPrice struct {
	productID int
	price     int
}

func productPriceAtDate(products [][]int) []productPrice {
	// products[i] = [product_id, new_price, change_date]
	// Find price of each product on 2019-08-16

	// Group by product
  // Membuat map untuk pencarian O(1): key → value
	changes := make(map[int][][2]int) // productID -> [(date, price)]
	for _, p := range products {
		id, price, date := p[0], p[1], p[2]
		changes[id] = append(changes[id], [2]int{date, price})
	}

	// Sort each product's changes by date
	for id := range changes {
		sort.Slice(changes[id], func(i, j int) bool {
			return changes[id][i][0] < changes[id][j][0]
		})
	}

	targetDate := 20190816
  // Membuat slice untuk menyimpan hasil
	result := make([]productPrice, 0, len(changes))

	for id, vals := range changes {
		price := 10 // default price
		for _, v := range vals {
			if v[0] <= targetDate {
				price = v[1]
			} else {
				break
			}
		}
		result = append(result, productPrice{id, price})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].productID < result[j].productID
	})

	return result
}

func main() {
	products := [][]int{
		{1, 20, 20190801},
		{2, 50, 20190801},
		{1, 30, 20190815},
		{1, 40, 20190817},
		{2, 80, 20190814},
	}
	fmt.Printf("%v (expected: [{1 30} {2 80}])\n", productPriceAtDate(products))

	products2 := [][]int{
		{1, 20, 20190817},
		{2, 50, 20190817},
	}
	fmt.Printf("%v (expected: [{1 10} {2 10}])\n", productPriceAtDate(products2))
}
```
