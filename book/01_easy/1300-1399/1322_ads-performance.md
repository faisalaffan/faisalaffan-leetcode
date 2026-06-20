# 1322 — Ads Performance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func AdsPerformance() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1322: Ads Performance
// https://leetcode.com/problems/ads-performance/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Ads (ad_id, user_id, action) where action is enum('Clicked','Viewed','Ignored')

import "fmt"

func main() {
	fmt.Println(AdsPerformance())
}

// Time: N/A (SQL query), Space: N/A
func AdsPerformance() string {
	return `SELECT
  ad_id,
  ROUND(IFNULL(SUM(action = 'Clicked') / SUM(action IN ('Clicked', 'Viewed')) * 100, 0), 2) AS ctr
FROM Ads
GROUP BY ad_id
ORDER BY ctr DESC, ad_id ASC;`
}
```
