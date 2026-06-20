# 1613 — Find The Missing Ids

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindMissingIDs(customerIDs []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(N log N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1613: Find the Missing IDs
// https://leetcode.com/problems/find-the-missing-ids/
// Difficulty: Medium [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	// SQL problem: find missing customer IDs within the range.
	// Translated to Go.
	// Table: Customers(customer_id)

	customerIDs := []int{1, 2, 4, 7, 8, 10}
	missing := FindMissingIDs(customerIDs)
	fmt.Println("Missing IDs:", missing)
}

func FindMissingIDs(customerIDs []int) []int {
	// Time: O(N log N), Space: O(1)
	if len(customerIDs) == 0 {
		return nil
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(customerIDs)
  // Alokasi slice integer
	result := make([]int, 0)

	// IDs range from 1 to max(customerID)
	for i := 1; i < customerIDs[len(customerIDs)-1]; i++ {
		// Binary search
		idx := sort.SearchInts(customerIDs, i)
		if idx == len(customerIDs) || customerIDs[idx] != i {
			result = append(result, i)
		}
	}

	return result
}
```
