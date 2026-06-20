# 1148 — Article Views I

## Deskripsi

**Soal:** [1148. Article Views I](https://leetcode.com/problems/article-views-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1148: Article Views I
// https://leetcode.com/problems/article-views-i/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT DISTINCT author_id AS id FROM Views WHERE author_id = viewer_id ORDER BY author_id")
}

// This is a SQL problem. The answer is the SQL query above.
```
