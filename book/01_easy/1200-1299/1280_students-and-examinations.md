# 1280 — Students And Examinations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1280: Students and Examinations
// https://leetcode.com/problems/students-and-examinations/
// Difficulty: Easy (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT s.student_id, s.student_name, sub.subject_name, COUNT(e.subject_name) AS attended_exams FROM Students s CROSS JOIN Subjects sub LEFT JOIN Examinations e ON s.student_id = e.student_id AND sub.subject_name = e.subject_name GROUP BY s.student_id, s.student_name, sub.subject_name ORDER BY s.student_id, sub.subject_name")
}

// This is a SQL problem. The answer is the SQL query above.
```
