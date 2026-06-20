# 1179 — Reformat Department Table

## Deskripsi

**Soal:** [1179. Reformat Department Table](https://leetcode.com/problems/reformat-department-table/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1179: Reformat Department Table
// https://leetcode.com/problems/reformat-department-table/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT id, SUM(CASE WHEN month='Jan' THEN revenue ELSE NULL END) AS Jan_Revenue, SUM(CASE WHEN month='Feb' THEN revenue ELSE NULL END) AS Feb_Revenue, SUM(CASE WHEN month='Mar' THEN revenue ELSE NULL END) AS Mar_Revenue, SUM(CASE WHEN month='Apr' THEN revenue ELSE NULL END) AS Apr_Revenue, SUM(CASE WHEN month='May' THEN revenue ELSE NULL END) AS May_Revenue, SUM(CASE WHEN month='Jun' THEN revenue ELSE NULL END) AS Jun_Revenue, SUM(CASE WHEN month='Jul' THEN revenue ELSE NULL END) AS Jul_Revenue, SUM(CASE WHEN month='Aug' THEN revenue ELSE NULL END) AS Aug_Revenue, SUM(CASE WHEN month='Sep' THEN revenue ELSE NULL END) AS Sep_Revenue, SUM(CASE WHEN month='Oct' THEN revenue ELSE NULL END) AS Oct_Revenue, SUM(CASE WHEN month='Nov' THEN revenue ELSE NULL END) AS Nov_Revenue, SUM(CASE WHEN month='Dec' THEN revenue ELSE NULL END) AS Dec_Revenue FROM Department GROUP BY id")
}

// This is a SQL problem. The answer is the SQL query above.
```
