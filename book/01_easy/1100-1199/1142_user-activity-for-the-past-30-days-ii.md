# 1142 — User Activity For The Past 30 Days Ii

## Deskripsi

**Soal:** [1142. User Activity For The Past 30 Days Ii](https://leetcode.com/problems/user-activity-for-the-past-30-days-ii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

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
