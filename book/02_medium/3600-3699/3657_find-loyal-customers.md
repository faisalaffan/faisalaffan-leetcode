# 3657 — Find Loyal Customers

## Deskripsi

**Soal:** [3657. Find Loyal Customers](https://leetcode.com/problems/find-loyal-customers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func findLoyalCustomers(purchases [][]int, minPurchases int, minAmount float64) []int`

## Solusi Go

```go
package main

// LeetCode #3657: Find Loyal Customers
// https://leetcode.com/problems/find-loyal-customers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findLoyalCustomers(purchases [][]int, minPurchases int, minAmount float64) []int {
  // Membuat map untuk pencarian O(1): key → value
	customerTotals := make(map[int]int)
  // Membuat map untuk pencarian O(1): key → value
	customerCounts := make(map[int]int)

	for _, p := range purchases {
		custID, amount := p[0], p[1]
		customerTotals[custID] += amount
		customerCounts[custID]++
	}

	var loyal []int
	for id := range customerTotals {
		if customerCounts[id] >= minPurchases && float64(customerTotals[id]) >= minAmount {
			loyal = append(loyal, id)
		}
	}

	sort.Ints(loyal)
	return loyal
}

func main() {
	fmt.Println(findLoyalCustomers([][]int{{1, 100}, {2, 50}, {1, 200}, {3, 300}, {2, 150}, {1, 50}}, 2, 200))
	fmt.Println(findLoyalCustomers([][]int{{1, 50}, {1, 50}}, 2, 100))
	fmt.Println(findLoyalCustomers([][]int{{1, 100}, {2, 200}}, 2, 100))
}
```
