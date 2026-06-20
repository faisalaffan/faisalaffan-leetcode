# 1831 — Maximum Transaction Each Day

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func maxTransactionPerDay(transactions []Transaction) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

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
  // HashMap: O(1) lookup
	dayMax := make(map[int]int)
	for _, t := range transactions {
		if t.Amount > dayMax[t.Day] {
			dayMax[t.Day] = t.Amount
		}
	}

	// Find transaction IDs that have max amount for their day
  // Alokasi slice
	result := make([]int, 0)
	for _, t := range transactions {
		if t.Amount == dayMax[t.Day] {
			result = append(result, t.ID)
		}
	}
  // Sort O(n log n)
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
