# 1565 — Unique Orders And Customers Per Month

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func UniqueOrdersAndCustomersPerMonth() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


## 💻 Solusi Go

```go
package main

// LeetCode #1565: Unique Orders and Customers Per Month
// https://leetcode.com/problems/unique-orders-and-customers-per-month/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Orders (order_id, order_date, customer_id, invoice)

import "fmt"

func main() {
	fmt.Println(UniqueOrdersAndCustomersPerMonth())
}

// Time: N/A (SQL query), Space: N/A
func UniqueOrdersAndCustomersPerMonth() string {
	return `SELECT
  DATE_FORMAT(order_date, '%Y-%m') AS month,
  COUNT(DISTINCT order_id) AS order_count,
  COUNT(DISTINCT customer_id) AS customer_count
FROM Orders
WHERE invoice > 20
GROUP BY month
ORDER BY month;`
}
```
