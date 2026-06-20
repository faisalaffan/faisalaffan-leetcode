# 1571 — Warehouse Manager

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func WarehouseManager() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
