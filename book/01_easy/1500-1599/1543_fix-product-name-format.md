# 1543 — Fix Product Name Format

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func FixProductNameFormat() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

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
