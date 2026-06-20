# 1303 — Find The Team Size

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1303: Find the Team Size
// https://leetcode.com/problems/find-the-team-size/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT employee_id, COUNT(team_id) OVER (PARTITION BY team_id) AS team_size FROM Employee")
}

// This is a SQL problem. The answer is the SQL query above.
```
