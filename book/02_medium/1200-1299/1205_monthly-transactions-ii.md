# 1205 — Monthly Transactions Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func monthlyTransactionsII(transactions []tx) []monthlyStat`

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

// LeetCode #1205: Monthly Transactions II
// https://leetcode.com/problems/monthly-transactions-ii/
// Difficulty: Medium [Paid]

// For each month and country: approved_count, approved_amount,
// chargeback_count, chargeback_amount.

// Time: O(n log n)
// Space: O(n)

type tx struct {
	id          int
	country     string
	state       string
	amount      int
	transDate   string
	chargebackDate string
}

type monthlyStat struct {
	month            string
	country          string
	approvedCount    int
	approvedAmount   int
	chargebackCount  int
	chargebackAmount int
}

func monthlyTransactionsII(transactions []tx) []monthlyStat {
	type key struct {
		month   string
		country string
	}
  // HashMap: O(1) lookup
	stats := make(map[key]monthlyStat)

	for _, t := range transactions {
		if t.state == "approved" || t.state == "declined" {
			month := t.transDate[:7]
			k := key{month, t.country}
			s := stats[k]
			s.month = month
			s.country = t.country
			if t.state == "approved" {
				s.approvedCount++
				s.approvedAmount += t.amount
			}
			stats[k] = s
		}
		if t.chargebackDate != "" {
			month := t.chargebackDate[:7]
			k := key{month, t.country}
			s := stats[k]
			s.month = month
			s.country = t.country
			s.chargebackCount++
			s.chargebackAmount += t.amount
			stats[k] = s
		}
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
	txs := []tx{
		{1, "US", "approved", 1000, "2018-12-01", ""},
		{2, "US", "declined", 2000, "2018-12-02", ""},
		{3, "US", "approved", 2000, "2019-01-01", "2019-01-05"},
	}
	result := monthlyTransactionsII(txs)
	for _, r := range result {
		fmt.Printf("month=%s country=%s approved=%d/%d chargeback=%d/%d\n",
			r.month, r.country, r.approvedCount, r.approvedAmount, r.chargebackCount, r.chargebackAmount)
	}
}
```
