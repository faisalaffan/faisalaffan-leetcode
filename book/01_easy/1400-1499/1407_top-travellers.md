# 1407 — Top Travellers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func TopTravellers() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


## 💻 Solusi Go

```go
package main

// LeetCode #1407: Top Travellers
// https://leetcode.com/problems/top-travellers/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Users (id, name), Rides (id, user_id, distance)

import "fmt"

func main() {
	fmt.Println(TopTravellers())
}

// Time: N/A (SQL query), Space: N/A
func TopTravellers() string {
	return `SELECT u.name, IFNULL(SUM(r.distance), 0) AS travelled_distance
FROM Users u
LEFT JOIN Rides r ON u.id = r.user_id
GROUP BY u.id, u.name
ORDER BY travelled_distance DESC, u.name ASC;`
}
```
