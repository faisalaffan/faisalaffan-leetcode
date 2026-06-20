# 2986 — Find Third Transaction

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findThirdTransaction(transactions []Transaction) []ThirdTransactionResult
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2986: Find Third Transaction
// https://leetcode.com/problems/find-third-transaction/
// Difficulty: Medium (SQL problem — simulated in Go)
//
// Simulates: For each user with >=3 transactions, find the 3rd transaction
// (by date ASC). Include only if the 3rd transaction's spend > BOTH the
// 1st and 2nd transaction's spend.

import (
	"fmt"
	"sort"
)

// Transaction represents the Transactions database table.
type Transaction struct {
	UserID          int
	Spend           float64
	TransactionDate string
}

// ThirdTransactionResult holds the output for each qualifying user.
type ThirdTransactionResult struct {
	UserID                int
	ThirdTransactionSpend float64
	ThirdTransactionDate  string
}

// findThirdTransaction simulates the SQL query.
// Time: O(n log n) | Space: O(n)
// n = number of transactions per user, overall O(n log n) due to sorting.
func findThirdTransaction(transactions []Transaction) []ThirdTransactionResult {
	// Group transactions by user_id.
  // Membuat map (HashMap) — pencarian O(1)
	userTxns := make(map[int][]Transaction)
	for _, t := range transactions {
		userTxns[t.UserID] = append(userTxns[t.UserID], t)
	}

	var results []ThirdTransactionResult

	for _, txns := range userTxns {
		if len(txns) < 3 {
			continue
		}
		// Sort by transaction_date ASC.
  // Custom sort dengan comparator
		sort.Slice(txns, func(i, j int) bool {
			return txns[i].TransactionDate < txns[j].TransactionDate
		})

		first := txns[0].Spend
		second := txns[1].Spend
		third := txns[2].Spend

		// Check if 3rd transaction's spend > both 1st and 2nd.
		if third > first && third > second {
			results = append(results, ThirdTransactionResult{
				UserID:                txns[0].UserID,
				ThirdTransactionSpend: third,
				ThirdTransactionDate:  txns[2].TransactionDate,
			})
		}
	}

	// Order by user_id ASC.
  // Custom sort dengan comparator
	sort.Slice(results, func(i, j int) bool {
		return results[i].UserID < results[j].UserID
	})

	return results
}

func main() {
	// Test data from the problem.
	transactions := []Transaction{
		{UserID: 1, Spend: 7.44, TransactionDate: "2022-07-11"},
		{UserID: 1, Spend: 49.78, TransactionDate: "2022-07-12"},
		{UserID: 1, Spend: 65.56, TransactionDate: "2022-07-13"},
		{UserID: 1, Spend: 30.00, TransactionDate: "2022-07-14"},
		// User 2 has 3 transactions, but 3rd (25.00) is not > 2nd (30.00).
		{UserID: 2, Spend: 10.00, TransactionDate: "2022-07-11"},
		{UserID: 2, Spend: 30.00, TransactionDate: "2022-07-12"},
		{UserID: 2, Spend: 25.00, TransactionDate: "2022-07-13"},
		// User 3 has only 2 transactions, does not qualify.
		{UserID: 3, Spend: 5.00, TransactionDate: "2022-07-11"},
		{UserID: 3, Spend: 8.00, TransactionDate: "2022-07-12"},
	}

	results := findThirdTransaction(transactions)

	fmt.Println("Third Transaction Results (user_id | third_transaction_spend | third_transaction_date):")
	for _, r := range results {
		fmt.Printf("%d | %.2f | %s\n", r.UserID, r.ThirdTransactionSpend, r.ThirdTransactionDate)
	}
	// Expected output:
	// 1 | 65.56 | 2022-07-13
}
```
