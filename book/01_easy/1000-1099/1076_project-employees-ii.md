# 1076 — Project Employees Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1076: Project Employees II
// https://leetcode.com/problems/project-employees-ii/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT project_id FROM Project GROUP BY project_id HAVING COUNT(employee_id) = (SELECT COUNT(employee_id) FROM Project GROUP BY project_id ORDER BY COUNT(employee_id) DESC LIMIT 1)")
}

// This is a SQL problem. The answer is the SQL query above.
```
