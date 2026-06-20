# 1495 — Friendly Movies Streamed Last Month

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func FriendlyMoviesStreamedLastMonth() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


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
