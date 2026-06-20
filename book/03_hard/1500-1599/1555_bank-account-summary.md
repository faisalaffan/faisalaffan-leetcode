# 1555 — Bank Account Summary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func accountSummary(users []User, transactions []Transaction) []resultRow1555`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1555: Bank Account Summary
// https://leetcode.com/problems/bank-account-summary/
// Difficulty: Hard
//
// Given users, transactions, and transfers, compute the net balance for each user
// and classify them as "Low Salary", "Average Salary", or "High Salary"
// based on their final balance (credit - debit + paid_by_incoming_transfers - paid_by_outgoing_transfers).
//
// Classification rules:
// - Low Salary:  balance < 20000
// - Average Salary: 20000 <= balance <= 50000
// - High Salary: balance > 50000

// User represents a bank user.
type User struct {
	UserID int
	Name   string
}

// Transaction is a credit/debit operation (credit means money IN to the user).
type Transaction struct {
	TransactionID int
	PaidBy        int
	PaidTo        int
	Amount        int
	TransactedOn  string
}

type resultRow1555 struct {
	Name    string
	Balance int
	Class   string
}

// accountSummary computes each user's balance and salary classification.
func accountSummary(users []User, transactions []Transaction) []resultRow1555 {
  // HashMap: O(1) lookup
	balances := make(map[int]int) // userID -> net balance

	// Initialize all users with 0
	for _, u := range users {
		balances[u.UserID] = 0
	}

	// Process transactions: paid_by loses money, paid_to gains money
	for _, t := range transactions {
		balances[t.PaidBy] -= t.Amount
		balances[t.PaidTo] += t.Amount
	}

	var results []resultRow1555

	for _, u := range users {
		bal := balances[u.UserID]
		class := ""
		if bal < 20000 {
			class = "Low Salary"
		} else if bal <= 50000 {
			class = "Average Salary"
		} else {
			class = "High Salary"
		}
		results = append(results, resultRow1555{
			Name:    u.Name,
			Balance: bal,
			Class:   class,
		})
	}

	// Sort by name ascending
  // Custom sort
	sort.Slice(results, func(i, j int) bool {
		return results[i].Name < results[j].Name
	})

	return results
}

func main() {
	users := []User{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}
	transactions := []Transaction{
		{1, 1, 2, 10000, "2023-01-01"},
		{2, 2, 3, 5000, "2023-01-02"},
		{3, 3, 1, 30000, "2023-01-03"},
		{4, 1, 2, 40000, "2023-01-04"},
	}

	results := accountSummary(users, transactions)
	for _, r := range results {
		fmt.Printf("%s | Balance: %d | %s\n", r.Name, r.Balance, r.Class)
	}
	// Expected:
	// Alice | Balance: -20000 | Low Salary (paid 10000+40000=50000, received 30000, net -20000)
	// Bob   | Balance: 45000  | Average Salary (paid 5000, received 10000+40000=50000, net 45000)
	// Charlie | Balance: -25000 | Low Salary (paid 30000, received 5000, net -25000)
}
```
