# 1407 — Top Travellers

## Deskripsi

**Soal:** [1407. Top Travellers](https://leetcode.com/problems/top-travellers/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

**Algoritma:** —

## Solusi Go

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
