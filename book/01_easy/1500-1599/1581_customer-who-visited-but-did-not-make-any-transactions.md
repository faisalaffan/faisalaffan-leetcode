# 1581 — Customer Who Visited But Did Not Make Any Transactions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func CustomerWhoVisitedButDidNotMakeAnyTransactions() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1581: Customer Who Visited but Did Not Make Any Transactions
// https://leetcode.com/problems/customer-who-visited-but-did-not-make-any-transactions/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Tables: Visits (visit_id, customer_id), Transactions (transaction_id, visit_id, amount)

import "fmt"

func main() {
	fmt.Println(CustomerWhoVisitedButDidNotMakeAnyTransactions())
}

// Time: N/A (SQL query), Space: N/A
func CustomerWhoVisitedButDidNotMakeAnyTransactions() string {
	return `SELECT v.customer_id, COUNT(*) AS count_no_trans
FROM Visits v
LEFT JOIN Transactions t ON v.visit_id = t.visit_id
WHERE t.visit_id IS NULL
GROUP BY v.customer_id;`
}
```
