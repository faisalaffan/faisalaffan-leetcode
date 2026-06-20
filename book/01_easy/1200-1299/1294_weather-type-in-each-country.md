# 1294 — Weather Type In Each Country

## Deskripsi

**Soal:** [1294. Weather Type In Each Country](https://leetcode.com/problems/weather-type-in-each-country/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Trie (pohon awalan)

## Solusi Go

```go
package main

// LeetCode #1294: Weather Type in Each Country
// https://leetcode.com/problems/weather-type-in-each-country/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT c.country_name, CASE WHEN AVG(w.weather_state * 1.0) <= 15 THEN 'Cold' WHEN AVG(w.weather_state * 1.0) >= 25 THEN 'Hot' ELSE 'Warm' END AS weather_type FROM Countries c JOIN Weather w ON c.country_id = w.country_id WHERE w.day BETWEEN '2019-11-01' AND '2019-11-30' GROUP BY c.country_id, c.country_name")
}

// This is a SQL problem. The answer is the SQL query above.
```
