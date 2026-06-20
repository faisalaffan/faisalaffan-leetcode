# 1421 — Npv Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func NpvQueries() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


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
