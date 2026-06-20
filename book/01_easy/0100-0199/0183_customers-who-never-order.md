# 0183 — Customers Who Never Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CustomersWhoNeverOrder() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #183: Customers Who Never Order
// https://leetcode.com/problems/customers-who-never-order/
// Difficulty: Easy

import "fmt"

func CustomersWhoNeverOrder() string {
	return "SELECT c.name AS Customers FROM Customers c LEFT JOIN Orders o ON c.id = o.customerId WHERE o.customerId IS NULL"
}

func main() {
	fmt.Println(CustomersWhoNeverOrder())
}
```
