# 1336 — Number Of Transactions Per Visit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func countTransactionsPerVisit(visits []visit, transactions []transaction) []rowCount`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1336: Number of Transactions per Visit
// https://leetcode.com/problems/number-of-transactions-per-visit/
// Difficulty: Hard [Paid]
//
// Approach: Simulate SQL aggregation in Go.
// Given a list of visits (user_id, visit_date) and transactions
// (user_id, visit_date, amount), count the distribution of
// number of transactions per visit. Return (transactions_count, visits_count)
// for each count from 0 up to the maximum transactions per visit.
//
// A visit with zero transactions (no matching transaction row) is also counted.

import (
	"fmt"
	"sort"
)

type visit struct {
	userID int
	date   string
}

type transaction struct {
	userID int
	date   string
	amount int
}

type rowCount struct {
	txCount int
	visits  int
}

func countTransactionsPerVisit(visits []visit, transactions []transaction) []rowCount {
	// Count transactions per visit
  // HashMap: O(1) lookup
	txPerVisit := make(map[[2]string]int)
	for _, t := range transactions {
		key := [2]string{fmt.Sprintf("%d", t.userID), t.date}
		txPerVisit[key]++
	}

	// Count visits per transaction count
  // HashMap: O(1) lookup
	visitDist := make(map[int]int)
	for _, v := range visits {
		key := [2]string{fmt.Sprintf("%d", v.userID), v.date}
		cnt := txPerVisit[key]
		visitDist[cnt]++
	}

	// Build result with all counts from 0..max
	maxCnt := 0
	for c := range visitDist {
		if c > maxCnt {
			maxCnt = c
		}
	}

	res := make([]rowCount, 0, maxCnt+1)
	for c := 0; c <= maxCnt; c++ {
		if v, ok := visitDist[c]; ok {
			res = append(res, rowCount{c, v})
		} else {
			res = append(res, rowCount{c, 0})
		}
	}

  // Custom sort
	sort.Slice(res, func(i, j int) bool { return res[i].txCount < res[j].txCount })
	return res
}

func main() {
	visits := []visit{
		{1, "2020-01-01"},
		{2, "2020-01-01"},
		{1, "2020-01-02"},
		{2, "2020-01-02"},
		{2, "2020-01-03"},
	}
	transactions := []transaction{
		{1, "2020-01-01", 100},
		{1, "2020-01-01", 200},
		{2, "2020-01-01", 50},
		{1, "2020-01-02", 150},
	}

	result := countTransactionsPerVisit(visits, transactions)
	for _, r := range result {
		fmt.Printf("transactions_count=%d, visits_count=%d\n", r.txCount, r.visits)
	}
}
```
