# 1587 — Bank Account Summary Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func BankAccountSummaryIi() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


## 💻 Solusi Go

```go
package main

// LeetCode #1587: Bank Account Summary II
// https://leetcode.com/problems/bank-account-summary-ii/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Users (account, name), Transactions (trans_id, account, amount, transacted_on)

import "fmt"

func main() {
	fmt.Println(BankAccountSummaryIi())
}

// Time: N/A (SQL query), Space: N/A
func BankAccountSummaryIi() string {
	return `SELECT u.name, SUM(t.amount) AS balance
FROM Users u
JOIN Transactions t ON u.account = t.account
GROUP BY u.account, u.name
HAVING SUM(t.amount) > 10000;`
}
```
