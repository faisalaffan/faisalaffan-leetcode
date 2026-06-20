# 1511 — Customer Order Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func CustomerOrderFrequency() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1511: Customer Order Frequency
// https://leetcode.com/problems/customer-order-frequency/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Customers (customer_id, name, country), Product (product_id, description, price), Orders (order_id, customer_id, product_id, order_date, quantity)

import "fmt"

func main() {
	fmt.Println(CustomerOrderFrequency())
}

// Time: N/A (SQL query), Space: N/A
func CustomerOrderFrequency() string {
	return `SELECT c.customer_id, c.name
FROM Customers c
JOIN Orders o ON c.customer_id = o.customer_id
JOIN Product p ON o.product_id = p.product_id
WHERE o.order_date BETWEEN '2020-06-01' AND '2020-07-31'
GROUP BY c.customer_id, c.name
HAVING SUM(CASE WHEN o.order_date BETWEEN '2020-06-01' AND '2020-06-30' THEN o.quantity * p.price END) >= 100
   AND SUM(CASE WHEN o.order_date BETWEEN '2020-07-01' AND '2020-07-31' THEN o.quantity * p.price END) >= 100;`
}
```
