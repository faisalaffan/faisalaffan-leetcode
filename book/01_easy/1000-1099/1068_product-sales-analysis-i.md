# 1068 — Product Sales Analysis I

## Deskripsi

**Soal:** [1068. Product Sales Analysis I](https://leetcode.com/problems/product-sales-analysis-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1068: Product Sales Analysis I
// https://leetcode.com/problems/product-sales-analysis-i/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)
// Note: This is a SQL problem. Go implementation simulates the query logic.

import "fmt"

type Sale struct {
	saleID    int
	productID int
	year      int
	quantity  int
	price     int
}

type Product struct {
	productID   int
	productName string
}

func main() {
	sales := []Sale{
		{1, 100, 2008, 10, 5000},
		{2, 100, 2009, 12, 5000},
		{7, 200, 2011, 15, 9000},
	}
	products := []Product{
		{100, "Nokia"},
		{200, "Apple"},
		{300, "Samsung"},
	}
	// SQL equivalent: SELECT p.product_name, s.year, s.price
	// FROM Sales s JOIN Product p ON s.product_id = p.product_id
	fmt.Println(productSalesAnalysisI(sales, products))
}

// LeetCode submission: productSalesAnalysisI (SQL equivalent)
func productSalesAnalysisI(sales []Sale, products []Product) []map[string]int {
  // Membuat map untuk pencarian O(1): key → value
	prodMap := make(map[int]string)
	for _, p := range products {
		prodMap[p.productID] = p.productName
	}
	type result struct {
		name  string
		year  int
		price int
	}
	var ans []result
	for _, s := range sales {
		if name, ok := prodMap[s.productID]; ok {
			ans = append(ans, result{name, s.year, s.price})
		}
	}
	return nil // placeholder - SQL query is the real answer
}
```
