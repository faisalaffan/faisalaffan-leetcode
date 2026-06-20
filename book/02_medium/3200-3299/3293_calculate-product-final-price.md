# 3293 — Calculate Product Final Price

## Deskripsi

**Soal:** [3293. Calculate Product Final Price](https://leetcode.com/problems/calculate-product-final-price/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(p + d) Space: O(p)  
**Kompleksitas Ruang:** O(p)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3293: Calculate Product Final Price
// https://leetcode.com/problems/calculate-product-final-price/
// Difficulty: Medium
// Time: O(p + d) Space: O(p)

import (
	"fmt"
	"sort"
)

func main() {
	products := []Product{
		{1, "Electronics", 1000},
		{2, "Clothing", 50},
		{3, "Electronics", 1200},
		{4, "Home", 500},
	}
	discounts := []Discount{
		{"Electronics", 10},
		{"Clothing", 20},
	}
	fmt.Println(calculateFinalPrice(products, discounts))
	// Expected: [{1 900 Electronics} {2 40 Clothing} {3 1080 Electronics} {4 500 Home}]

	// Test 2: No discounts
	fmt.Println(calculateFinalPrice(
		[]Product{{1, "Food", 100}},
		[]Discount{{"Electronics", 10}},
	))
	// Expected: [{1 100 Food}]
}

type Product struct {
	ID       int
	Category string
	Price    float64
}

type Discount struct {
	Category string
	Percent  int
}

type ProductResult struct {
	ID         int
	FinalPrice float64
	Category   string
}

func calculateFinalPrice(products []Product, discounts []Discount) []ProductResult {
  // Membuat map untuk pencarian O(1): key → value
	discMap := make(map[string]int)
	for _, d := range discounts {
		discMap[d.Category] = d.Percent
	}

  // Membuat slice untuk menyimpan hasil
	res := make([]ProductResult, len(products))
	for i, p := range products {
		fp := p.Price
		if perc, ok := discMap[p.Category]; ok {
			fp = p.Price * (1.0 - float64(perc)/100.0)
		}
		res[i] = ProductResult{p.ID, fp, p.Category}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	return res
}
```
