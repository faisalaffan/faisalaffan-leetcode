# 1571 — Warehouse Manager

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func WarehouseManager() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


## 💻 Solusi Go

```go
package main

// LeetCode #1571: Warehouse Manager
// https://leetcode.com/problems/warehouse-manager/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Warehouse (name, product_id, units), Products (product_id, product_name, Width, Length, Height)

import "fmt"

func main() {
	fmt.Println(WarehouseManager())
}

// Time: N/A (SQL query), Space: N/A
func WarehouseManager() string {
	return `SELECT w.name AS warehouse_name, SUM(w.units * p.Width * p.Length * p.Height) AS volume
FROM Warehouse w
JOIN Products p ON w.product_id = p.product_id
GROUP BY w.name
ORDER BY warehouse_name;`
}
```
