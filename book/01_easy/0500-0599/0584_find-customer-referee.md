# 0584 — Find Customer Referee

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindCustomerReferee() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #584: Find Customer Referee
// https://leetcode.com/problems/find-customer-referee/
// Difficulty: Easy

import "fmt"

func FindCustomerReferee() string {
	return "SELECT name FROM Customer WHERE referee_id != 2 OR referee_id IS NULL"
}

func main() {
	fmt.Println(FindCustomerReferee())
}
```
