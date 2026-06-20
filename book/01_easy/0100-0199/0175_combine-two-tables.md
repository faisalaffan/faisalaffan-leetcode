# 0175 — Combine Two Tables

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func CombineTwoTables() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


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
