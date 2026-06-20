# 1435 — Create A Session Bar Chart

## Deskripsi

**Soal:** [1435. Create A Session Bar Chart](https://leetcode.com/problems/create-a-session-bar-chart/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

**Algoritma:** —

## Solusi Go

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
