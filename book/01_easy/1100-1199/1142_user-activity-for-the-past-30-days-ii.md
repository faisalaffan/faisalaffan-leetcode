# 1142 — User Activity For The Past 30 Days Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1142: User Activity for the Past 30 Days II
// https://leetcode.com/problems/user-activity-for-the-past-30-days-ii/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT IFNULL(ROUND(COUNT(DISTINCT session_id) / COUNT(DISTINCT user_id), 2), 0) AS average_sessions_per_user FROM Activity WHERE activity_date BETWEEN '2019-06-28' AND '2019-07-27'")
}

// This is a SQL problem. The answer is the SQL query above.
```
