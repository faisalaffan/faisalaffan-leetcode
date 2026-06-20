# 1205 — Monthly Transactions Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func monthlyTransactionsII(transactions []tx) []monthlyStat
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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
  // Custom sort dengan comparator
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
