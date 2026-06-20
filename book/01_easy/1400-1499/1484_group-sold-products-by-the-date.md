# 1484 — Group Sold Products By The Date

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func GroupSoldProductsByTheDate() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


## 💻 Solusi Go

```go
package main

// LeetCode #1484: Group Sold Products By The Date
// https://leetcode.com/problems/group-sold-products-by-the-date/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Activities (sell_date, product)

import "fmt"

func main() {
	fmt.Println(GroupSoldProductsByTheDate())
}

// Time: N/A (SQL query), Space: N/A
func GroupSoldProductsByTheDate() string {
	return `SELECT
  sell_date,
  COUNT(DISTINCT product) AS num_sold,
  GROUP_CONCAT(DISTINCT product ORDER BY product SEPARATOR ',') AS products
FROM Activities
GROUP BY sell_date
ORDER BY sell_date;`
}
```
