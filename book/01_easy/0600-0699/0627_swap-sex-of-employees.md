# 0627 — Swap Sex Of Employees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SwapSexOfEmployees() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #627: Swap Sex of Employees
// https://leetcode.com/problems/swap-sex-of-employees/
// Difficulty: Easy

import "fmt"

func SwapSexOfEmployees() string {
	return "UPDATE Salary SET sex = CASE WHEN sex = 'm' THEN 'f' ELSE 'm' END"
}

func main() {
	fmt.Println(SwapSexOfEmployees())
}
```
