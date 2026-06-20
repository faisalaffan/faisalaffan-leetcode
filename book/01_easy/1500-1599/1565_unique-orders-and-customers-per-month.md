# 1565 — Unique Orders And Customers Per Month

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func UniqueOrdersAndCustomersPerMonth() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
