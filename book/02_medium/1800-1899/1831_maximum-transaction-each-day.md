# 1831 — Maximum Transaction Each Day

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxTransactionPerDay(transactions []Transaction) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1831: Maximum Transaction Each Day
// https://leetcode.com/problems/maximum-transaction-each-day/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

type Transaction struct {
	ID     int
	Day    int
	Amount int
}

func maxTransactionPerDay(transactions []Transaction) []int {
	// Group by day, find max amount
  // Membuat map (HashMap) — pencarian O(1)
	dayMax := make(map[int]int)
	for _, t := range transactions {
		if t.Amount > dayMax[t.Day] {
			dayMax[t.Day] = t.Amount
		}
	}

	// Find transaction IDs that have max amount for their day
  // Alokasi slice integer
	result := make([]int, 0)
	for _, t := range transactions {
		if t.Amount == dayMax[t.Day] {
			result = append(result, t.ID)
		}
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	transactions := []Transaction{
		{1, 1, 100},
		{2, 1, 200},
		{3, 2, 150},
		{4, 2, 100},
	}
	fmt.Println(maxTransactionPerDay(transactions)) // Expected: [2, 3]
}
```
