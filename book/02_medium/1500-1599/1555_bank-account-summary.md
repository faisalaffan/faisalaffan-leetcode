# 1555 — Bank Account Summary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func AccountSummary(users map[int]string, credits map[int]int, txns []struct{ userID, amount int }) []accountInfo`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1555: Bank Account Summary
// https://leetcode.com/problems/bank-account-summary/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// SQL problem: find users with balance > 10000 or overdrawn users.
	// Translated to Go.
	// Tables: Users(user_id, user_name, credit), Transactions(trans_id, user_id, amount, transacted_on)

	users := map[int]string{1: "Alice", 2: "Bob", 3: "Charlie"}
	credits := map[int]int{1: 5000, 2: 8000, 3: 3000}
	txns := []struct{ userID, amount int }{
		{1, 2000}, {2, -3000}, {3, 10000}, {1, -1000},
	}

	result := AccountSummary(users, credits, txns)
	fmt.Println("Account summary:")
	for _, r := range result {
		fmt.Printf("  %s: balance=%d, status=%s\n", r.name, r.balance, r.status)
	}
}

type accountInfo struct {
	userID  int
	name    string
	balance int
	status  string
}

func AccountSummary(users map[int]string, credits map[int]int, txns []struct{ userID, amount int }) []accountInfo {
  // HashMap: O(1) lookup
	balance := make(map[int]int)
	for uid, credit := range credits {
		balance[uid] = credit
	}
	for _, t := range txns {
		balance[t.userID] += t.amount
	}

	result := make([]accountInfo, 0)
	for uid, name := range users {
		b := balance[uid]
		status := "Normal"
		if b < 0 {
			status = "Overdrawn"
		}
		result = append(result, accountInfo{uid, name, b, status})
	}

	return result
}
```
