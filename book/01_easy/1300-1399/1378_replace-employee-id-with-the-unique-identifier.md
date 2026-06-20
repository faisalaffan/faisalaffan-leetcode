# 1378 — Replace Employee Id With The Unique Identifier

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func ReplaceEmployeeIdWithTheUniqueIdentifier() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
