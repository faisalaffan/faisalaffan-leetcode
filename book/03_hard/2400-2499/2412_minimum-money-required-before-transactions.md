# 2412 — Minimum Money Required Before Transactions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumMoney(transactions [][]int) int64
```

> **💡 Hint:** Split transactions into two groups:

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2412: Minimum Money Required Before Transactions
// https://leetcode.com/problems/minimum-money-required-before-transactions/
// Difficulty: Hard
//
// You are given a list of transactions where each transaction is [cost, cashback].
// You can execute transactions in any order. You start with some money, and for
// each transaction: if money >= cost, you pay cost and receive cashback (money
// = money - cost + cashback). Find the minimum starting money needed to complete
// all transactions.
//
// Approach: Split transactions into two groups:
//   - "loss" transactions: cost > cashback (net loss)
//   - "gain" transactions: cost <= cashback (net gain or break-even)
//
// For maximum safety, process loss transactions first (they deplete money).
// Within loss transactions, process those with the highest cashback first
// (to maximize money recovery). Within gain transactions, process those with
// the lowest cost first (to conserve money).
//
// The minimum starting money = max over all prefixes of (cumulative money needed).

import (
	"fmt"
	"sort"
)

func minimumMoney(transactions [][]int) int64 {
	// Separate loss and gain transactions
  // Membuat matriks/slice 2D untuk DP
	loss := make([][]int, 0)
  // Membuat matriks/slice 2D untuk DP
	gain := make([][]int, 0)

	for _, t := range transactions {
		cost, cashback := t[0], t[1]
		if cost > cashback {
			loss = append(loss, t)
		} else {
			gain = append(gain, t)
		}
	}

	// Sort loss transactions by cashback descending (recover more money sooner)
  // Custom sort dengan comparator
	sort.Slice(loss, func(i, j int) bool {
		return loss[i][1] > loss[j][1]
	})

	// Sort gain transactions by cost ascending (spend less money first)
  // Custom sort dengan comparator
	sort.Slice(gain, func(i, j int) bool {
		return gain[i][0] < gain[j][0]
	})

	// Simulate transactions in order: loss first, then gain
	ordered := append(loss, gain...)

	var money int64
	var minStart int64

	for _, t := range ordered {
		cost, cashback := int64(t[0]), int64(t[1])

		// If we don't have enough money, increase starting money
		if money < cost {
			needed := cost - money
			minStart += needed
			money += needed
		}

		money = money - cost + cashback
	}

	return minStart
}

func main() {
	// Example 1
	fmt.Println(minimumMoney([][]int{{2, 1}, {5, 0}, {4, 2}}))
	// Example 2
	fmt.Println(minimumMoney([][]int{{3, 0}, {0, 3}}))
	// Single loss transaction
	fmt.Println(minimumMoney([][]int{{10, 2}}))
	// Single gain transaction
	fmt.Println(minimumMoney([][]int{{5, 10}}))
	// All gains
	fmt.Println(minimumMoney([][]int{{1, 2}, {2, 3}, {3, 5}}))
	// All losses
	fmt.Println(minimumMoney([][]int{{10, 1}, {8, 2}, {5, 3}}))
}
```
