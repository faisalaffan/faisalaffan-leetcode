# 0195 — Tenth Line

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func TenthLine() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #195: Tenth Line
// https://leetcode.com/problems/tenth-line/
// Difficulty: Easy

import "fmt"

func TenthLine() string {
	return "sed -n '10p' file.txt"
}

func main() {
	fmt.Println(TenthLine())
}
```
