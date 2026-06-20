# 1587 — Bank Account Summary Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func BankAccountSummaryIi() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
