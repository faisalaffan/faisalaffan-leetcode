# 1098 — Unpopular Books

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func unpopularBooks(books []string, orders []struct { name string quantity int daysAgo int }) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + m) where n = books, m = orders  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1098: Unpopular Books
// https://leetcode.com/problems/unpopular-books/
// Difficulty: Medium
//
// Approach: Filter books ordered less than 10 times in the last year
// Time: O(n + m) where n = books, m = orders
// Space: O(n)

import "fmt"

func main() {
	// Books: (book_id, name)
	books := []string{"Book A", "Book B", "Book C", "Book D"}
	// Orders: (book_name, quantity, days_ago)
	orders := []struct {
		name     string
		quantity int
		daysAgo  int
	}{
		{"Book A", 5, 30},
		{"Book B", 15, 10},
		{"Book C", 8, 20},
	}
	fmt.Println(unpopularBooks(books, orders))
}

func unpopularBooks(books []string, orders []struct {
	name     string
	quantity int
	daysAgo  int
}) []string {
  // HashMap: O(1) lookup
	orderCount := make(map[string]int)
	for _, o := range orders {
		if o.daysAgo <= 365 {
			orderCount[o.name] += o.quantity
		}
	}

	result := make([]string, 0)
	for _, b := range books {
		count, exists := orderCount[b]
		if !exists || count < 10 {
			result = append(result, b)
		}
	}

	return result
}
```
