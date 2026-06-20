# 1495 — Friendly Movies Streamed Last Month

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func FriendlyMoviesStreamedLastMonth() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1495: Friendly Movies Streamed Last Month
// https://leetcode.com/problems/friendly-movies-streamed-last-month/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: TVProgram (program_date, content_id, channel), Content (content_id, title, Kids_content, content_type)

import "fmt"

func main() {
	fmt.Println(FriendlyMoviesStreamedLastMonth())
}

// Time: N/A (SQL query), Space: N/A
func FriendlyMoviesStreamedLastMonth() string {
	return `SELECT DISTINCT c.title
FROM TVProgram p
JOIN Content c ON p.content_id = c.content_id
WHERE c.Kids_content = 'Y'
  AND c.content_type = 'Movies'
  AND p.program_date BETWEEN '2020-06-01' AND '2020-06-30';`
}
```
