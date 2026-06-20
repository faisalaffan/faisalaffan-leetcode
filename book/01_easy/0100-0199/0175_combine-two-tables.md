# 0175 — Combine Two Tables

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func CombineTwoTables() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #175: Combine Two Tables
// https://leetcode.com/problems/combine-two-tables/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
// Write your MySQL query statement below
func CombineTwoTables() string {
	return "SELECT Person.firstName, Person.lastName, Address.city, Address.state FROM Person LEFT JOIN Address ON Person.personId = Address.personId"
}

func main() {
	fmt.Println(CombineTwoTables())
}
```
