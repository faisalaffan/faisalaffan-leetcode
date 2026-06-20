# 1084 — Sales Analysis Iii

## Deskripsi

**Soal:** [1084. Sales Analysis Iii](https://leetcode.com/problems/sales-analysis-iii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1084: Sales Analysis III
// https://leetcode.com/problems/sales-analysis-iii/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT p.product_id, p.product_name FROM Product p JOIN Sales s ON p.product_id = s.product_id GROUP BY p.product_id, p.product_name HAVING MIN(s.sale_date) >= '2019-01-01' AND MAX(s.sale_date) <= '2019-03-31'")
}

// This is a SQL problem. The answer is the SQL query above.
```
