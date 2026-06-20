# 1251 — Average Selling Price

## Deskripsi

**Soal:** [1251. Average Selling Price](https://leetcode.com/problems/average-selling-price/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1251: Average Selling Price
// https://leetcode.com/problems/average-selling-price/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT p.product_id, IFNULL(ROUND(SUM(u.units * p.price) / SUM(u.units), 2), 0) AS average_price FROM Prices p LEFT JOIN UnitsSold u ON p.product_id = u.product_id AND u.purchase_date BETWEEN p.start_date AND p.end_date GROUP BY p.product_id")
}

// This is a SQL problem. The answer is the SQL query above.
```
