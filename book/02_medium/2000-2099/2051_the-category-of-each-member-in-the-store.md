# 2051 — The Category Of Each Member In The Store

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func categorizeMembers(members []Member) map[int]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2051: The Category of Each Member in the Store
// https://leetcode.com/problems/the-category-of-each-member-in-the-store/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Member struct {
	ID       int
	VisitCnt int
	Spent    int
}

func categorizeMembers(members []Member) map[int]string {
  // Custom sort
	sort.Slice(members, func(i, j int) bool {
		return members[i].ID < members[j].ID
	})

	// Find max visits and max spent for normalization
	maxVisits := 0
	maxSpent := 0
	for _, m := range members {
		if m.VisitCnt > maxVisits {
			maxVisits = m.VisitCnt
		}
		if m.Spent > maxSpent {
			maxSpent = m.Spent
		}
	}

  // HashMap: O(1) lookup
	result := make(map[int]string)
	for _, m := range members {
		// Determine category: premium, gold, silver, bronze, basic
		visitRatio := float64(m.VisitCnt) / float64(maxVisits)
		spentRatio := float64(m.Spent) / float64(maxSpent)
		score := visitRatio + spentRatio

		var category string
		if score >= 1.5 {
			category = "premium"
		} else if score >= 1.0 {
			category = "gold"
		} else if score >= 0.5 {
			category = "silver"
		} else if score > 0 {
			category = "bronze"
		} else {
			category = "basic"
		}
		result[m.ID] = category
	}
	return result
}

func main() {
	// Test case 1
	members1 := []Member{
		{1, 10, 1000},
		{2, 5, 500},
		{3, 1, 50},
	}
	result1 := categorizeMembers(members1)
	fmt.Println("Test 1:")
	for id := 1; id <= 3; id++ {
		fmt.Printf("  Member %d: %s\n", id, result1[id])
	}

	// Test case 2
	members2 := []Member{
		{1, 0, 0},
	}
	result2 := categorizeMembers(members2)
	fmt.Println("Test 2: Member 1:", result2[1])
	// Expected: basic
}
```
