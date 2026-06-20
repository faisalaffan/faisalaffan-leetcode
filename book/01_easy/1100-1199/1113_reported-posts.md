# 1113 — Reported Posts

## Deskripsi

**Soal:** [1113. Reported Posts](https://leetcode.com/problems/reported-posts/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1113: Reported Posts
// https://leetcode.com/problems/reported-posts/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT extra AS report_reason, COUNT(DISTINCT post_id) AS report_count FROM Actions WHERE action = 'report' AND action_date = '2019-07-04' GROUP BY extra")
}

// This is a SQL problem. The answer is the SQL query above.
```
