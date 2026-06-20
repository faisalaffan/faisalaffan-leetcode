# 0196 — Delete Duplicate Emails

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func DeleteDuplicateEmails() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #196: Delete Duplicate Emails
// https://leetcode.com/problems/delete-duplicate-emails/
// Difficulty: Easy

import "fmt"

func DeleteDuplicateEmails() string {
	return "DELETE p1 FROM Person p1, Person p2 WHERE p1.email = p2.email AND p1.id > p2.id"
}

func main() {
	fmt.Println(DeleteDuplicateEmails())
}
```
