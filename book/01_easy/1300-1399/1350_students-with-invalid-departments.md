# 1350 — Students With Invalid Departments

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func StudentsWithInvalidDepartments() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1350: Students With Invalid Departments
// https://leetcode.com/problems/students-with-invalid-departments/
// Difficulty: Easy [Paid]
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Students (id, name, department_id), Departments (id, name)

import "fmt"

func main() {
	fmt.Println(StudentsWithInvalidDepartments())
}

// Time: N/A (SQL query), Space: N/A
func StudentsWithInvalidDepartments() string {
	return `SELECT s.id, s.name
FROM Students s
LEFT JOIN Departments d ON s.department_id = d.id
WHERE d.id IS NULL;`
}
```
