# 1378 — Replace Employee Id With The Unique Identifier

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func ReplaceEmployeeIdWithTheUniqueIdentifier() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


## 💻 Solusi Go

```go
package main

// LeetCode #1378: Replace Employee ID With The Unique Identifier
// https://leetcode.com/problems/replace-employee-id-with-the-unique-identifier/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Employees (id, name), EmployeeUNI (id, unique_id)

import "fmt"

func main() {
	fmt.Println(ReplaceEmployeeIdWithTheUniqueIdentifier())
}

// Time: N/A (SQL query), Space: N/A
func ReplaceEmployeeIdWithTheUniqueIdentifier() string {
	return `SELECT eu.unique_id, e.name
FROM Employees e
LEFT JOIN EmployeeUNI eu ON e.id = eu.id;`
}
```
