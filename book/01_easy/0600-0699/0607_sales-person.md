# 0607 — Sales Person

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SalesPerson() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #607: Sales Person
// https://leetcode.com/problems/sales-person/
// Difficulty: Easy

import "fmt"

func SalesPerson() string {
	return "SELECT s.name FROM SalesPerson s WHERE s.sales_id NOT IN (SELECT o.sales_id FROM Orders o JOIN Company c ON o.com_id = c.com_id WHERE c.name = 'RED')"
}

func main() {
	fmt.Println(SalesPerson())
}
```
