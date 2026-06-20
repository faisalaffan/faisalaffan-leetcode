# 1435 — Create A Session Bar Chart

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func CreateASessionBarChart() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1435: Create a Session Bar Chart
// https://leetcode.com/problems/create-a-session-bar-chart/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Sessions (session_id, duration)

import "fmt"

func main() {
	fmt.Println(CreateASessionBarChart())
}

// Time: N/A (SQL query), Space: N/A
func CreateASessionBarChart() string {
	return `SELECT
  CASE
    WHEN duration / 60 BETWEEN 0 AND 4 THEN '[0-5>'
    WHEN duration / 60 BETWEEN 5 AND 9 THEN '[5-10>'
    WHEN duration / 60 BETWEEN 10 AND 14 THEN '[10-15>'
    ELSE '15 or more'
  END AS bin,
  COUNT(*) AS total
FROM Sessions
GROUP BY bin;`
}
```
