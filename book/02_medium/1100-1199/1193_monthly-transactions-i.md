# 1193 — Monthly Transactions I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func monthlyTransactionsI(transactions []trans) []monthlyStat`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1193: Monthly Transactions I
// https://leetcode.com/problems/monthly-transactions-i/
// Difficulty: Medium

// Group transactions by country and month, count approved and total.

// Time: O(n log n)
// Space: O(n)

type trans struct {
	id      int
	country string
	state   string
	amount  int
	date    string
}

type monthlyStat struct {
	month         string
	country       string
	approvedCount int
	approvedTotal int
	totalCount    int
	totalAmount   int
}

func monthlyTransactionsI(transactions []trans) []monthlyStat {
	type key struct {
		month   string
		country string
	}
  // HashMap: O(1) lookup
	stats := make(map[key]monthlyStat)

	for _, t := range transactions {
		month := t.date[:7]
		k := key{month, t.country}
		s := stats[k]
		s.month = month
		s.country = t.country
		s.totalCount++
		s.totalAmount += t.amount
		if t.state == "approved" {
			s.approvedCount++
			s.approvedTotal += t.amount
		}
		stats[k] = s
	}

	result := make([]monthlyStat, 0, len(stats))
	for _, s := range stats {
		result = append(result, s)
	}

  // Custom sort
	sort.Slice(result, func(i, j int) bool {
		if result[i].month != result[j].month {
			return result[i].month < result[j].month
		}
		return result[i].country < result[j].country
	})

	return result
}

func main() {
	transactions := []trans{
		{1, "US", "approved", 1000, "2018-12-01"},
		{2, "US", "declined", 2000, "2018-12-02"},
		{3, "US", "approved", 3000, "2019-01-01"},
		{4, "DE", "approved", 4000, "2019-01-02"},
	}
	result := monthlyTransactionsI(transactions)
	for _, r := range result {
		fmt.Printf("%s/%s: approved=%d/%d total=%d/%d\n",
			r.month, r.country, r.approvedCount, r.approvedTotal, r.totalCount, r.totalAmount)
	}
}
```
