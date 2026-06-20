# 2701 — Consecutive Transactions With Increasing Amounts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func consecutiveIncreasingTransactions(txns []Transaction) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2701: Consecutive Transactions with Increasing Amounts
// https://leetcode.com/problems/consecutive-transactions-with-increasing-amounts/
// Difficulty: Hard [Paid]
//
// Find customers who have at least 3 consecutive transactions (ordered by date)
// with strictly increasing amounts. Return [customer_id, count_of_consecutive_sets].
//
// Equivalent SQL: Given Transactions(customer_id, transaction_date, amount),
// find customers where there are 3+ consecutive rows with increasing amounts,
// grouped by customer and ordered by date.

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1: customer 1 has 3 consecutive increasing transactions
	txns1 := []Transaction{
		{1, "2024-01-01", 100},
		{1, "2024-01-02", 200},
		{1, "2024-01-03", 300},
		{2, "2024-01-01", 100},
		{2, "2024-01-02", 50},
	}
	fmt.Println(consecutiveIncreasingTransactions(txns1))

	// Test case 2: multiple customers with sequences
	txns2 := []Transaction{
		{1, "2024-01-01", 10},
		{1, "2024-01-02", 20},
		{1, "2024-01-03", 30},
		{1, "2024-01-04", 5},
		{2, "2024-01-01", 5},
		{2, "2024-01-02", 10},
		{2, "2024-01-03", 15},
		{2, "2024-01-04", 25},
	}
	fmt.Println(consecutiveIncreasingTransactions(txns2))

	// Test case 3: no increasing sequences
	txns3 := []Transaction{
		{1, "2024-01-01", 100},
		{1, "2024-01-02", 90},
		{1, "2024-01-03", 80},
	}
	fmt.Println(consecutiveIncreasingTransactions(txns3))
}

// Transaction represents a customer transaction.
type Transaction struct {
	CustomerID int
	Date       string // ISO format "YYYY-MM-DD"
	Amount     int
}

// Result records a customer and count of consecutive increasing sequences.
type Result struct {
	CustomerID int
	Count      int // number of consecutive increasing sequences of length >= 3
}

// consecutiveIncreasingTransactions finds customers with at least 3 consecutive
// transactions where amounts strictly increase.
// Returns [][]int where each inner is [customer_id, count].
func consecutiveIncreasingTransactions(txns []Transaction) [][]int {
	if len(txns) == 0 {
		return [][]int{}
	}

	// Group by customer
  // Membuat map (HashMap) — pencarian O(1)
	byCustomer := make(map[int][]Transaction)
	for _, t := range txns {
		byCustomer[t.CustomerID] = append(byCustomer[t.CustomerID], t)
	}

	type result struct {
		customerID int
		count      int
	}
	var results []result

	for cid, txns := range byCustomer {
		// Sort by date for each customer
  // Custom sort dengan comparator
		sort.Slice(txns, func(i, j int) bool {
			return txns[i].Date < txns[j].Date
		})

		// Sliding window: find consecutive increasing runs
		n := len(txns)
		incLen := 1 // length of current increasing run
		totalSets := 0

		for i := 1; i < n; i++ {
			if txns[i].Amount > txns[i-1].Amount {
				incLen++
			} else {
				if incLen >= 3 {
					totalSets += incLen - 2
					// incLen-2 counts the number of length-3+ sub-sequences
					// within this run: e.g., run of 4 has 2 sets of 3 consecutive
				}
				incLen = 1
			}
		}
		if incLen >= 3 {
			totalSets += incLen - 2
		}

		if totalSets > 0 {
			results = append(results, result{cid, totalSets})
		}
	}

	// Sort results by customer ID
  // Custom sort dengan comparator
	sort.Slice(results, func(i, j int) bool {
		return results[i].customerID < results[j].customerID
	})

  // Membuat matriks/slice 2D untuk DP
	out := make([][]int, len(results))
	for i, r := range results {
		out[i] = []int{r.customerID, r.count}
	}
	return out
}
```
