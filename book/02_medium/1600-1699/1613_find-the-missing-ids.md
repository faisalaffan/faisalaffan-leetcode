# 1613 — Find The Missing Ids

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func FindMissingIDs(customerIDs []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Sorting

**Waktu:** O(N log N), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

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

  // Sort O(n log n)
	sort.Ints(customerIDs)
  // Alokasi slice
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
