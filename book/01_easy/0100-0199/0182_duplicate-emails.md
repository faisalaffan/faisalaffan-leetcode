# 0182 — Duplicate Emails

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func DuplicateEmails() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #182: Duplicate Emails
// https://leetcode.com/problems/duplicate-emails/
// Difficulty: Easy

import "fmt"

func DuplicateEmails() string {
	return "SELECT email FROM Person GROUP BY email HAVING COUNT(email) > 1"
}

func main() {
	fmt.Println(DuplicateEmails())
}
```
