# 1327 — List The Products Ordered In A Period

## Deskripsi

**Soal:** [1327. List The Products Ordered In A Period](https://leetcode.com/problems/list-the-products-ordered-in-a-period/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #1327: List the Products Ordered in a Period
// https://leetcode.com/problems/list-the-products-ordered-in-a-period/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Products (product_id, product_name, product_category), Orders (product_id, order_date, unit)

import "fmt"

func main() {
	fmt.Println(ListTheProductsOrderedInAPeriod())
}

// Time: N/A (SQL query), Space: N/A
func ListTheProductsOrderedInAPeriod() string {
	return `SELECT p.product_name, SUM(o.unit) AS unit
FROM Products p
JOIN Orders o ON p.product_id = o.product_id
WHERE o.order_date BETWEEN '2020-02-01' AND '2020-02-29'
GROUP BY p.product_id, p.product_name
HAVING SUM(o.unit) >= 100;`
}
```
