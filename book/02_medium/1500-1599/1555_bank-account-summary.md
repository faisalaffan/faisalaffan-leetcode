# 1555 — Bank Account Summary

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func AccountSummary(users map[int]string, credits map[int]int, txns []struct{ userID, amount int }) []accountInfo
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
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
