# 1421 — Npv Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func NpvQueries() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1421: NPV Queries
// https://leetcode.com/problems/npv-queries/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: NPV (id, year, npv), Queries (id, year)

import "fmt"

func main() {
	fmt.Println(NpvQueries())
}

// Time: N/A (SQL query), Space: N/A
func NpvQueries() string {
	return `SELECT q.id, q.year, IFNULL(n.npv, 0) AS npv
FROM Queries q
LEFT JOIN NPV n ON q.id = n.id AND q.year = n.year
ORDER BY q.id, q.year;`
}
```
