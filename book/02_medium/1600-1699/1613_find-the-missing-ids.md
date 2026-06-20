# 1613 — Find The Missing Ids

## Deskripsi

**Soal:** [1613. Find The Missing Ids](https://leetcode.com/problems/find-the-missing-ids/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N log N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Binary Search (pencarian biner)

## Solusi Go

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

	sort.Ints(customerIDs)
  // Membuat slice untuk menyimpan hasil
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
