# 1587 — Bank Account Summary Ii

## Deskripsi

**Soal:** [1587. Bank Account Summary Ii](https://leetcode.com/problems/bank-account-summary-ii/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

**Algoritma:** —

## Solusi Go

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
