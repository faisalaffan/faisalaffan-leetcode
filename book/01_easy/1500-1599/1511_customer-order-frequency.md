# 1511 — Customer Order Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func CustomerOrderFrequency() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


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
