# 1543 — Fix Product Name Format

## Deskripsi

**Soal:** [1543. Fix Product Name Format](https://leetcode.com/problems/fix-product-name-format/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1543: Fix Product Name Format
// https://leetcode.com/problems/fix-product-name-format/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Products (product_id, product_name, price)

import "fmt"

func main() {
	fmt.Println(FixProductNameFormat())
}

// Time: N/A (SQL query), Space: N/A
func FixProductNameFormat() string {
	return `SELECT
  TRIM(LOWER(product_name)) AS product_name,
  DATE_FORMAT(sale_date, '%Y-%m') AS sale_date,
  COUNT(*) AS total
FROM Sales
GROUP BY TRIM(LOWER(product_name)), DATE_FORMAT(sale_date, '%Y-%m')
ORDER BY product_name, sale_date;`
}
```
