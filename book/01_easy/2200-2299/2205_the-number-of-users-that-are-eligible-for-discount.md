# 2205 — The Number Of Users That Are Eligible For Discount

## Deskripsi

**Soal:** [2205. The Number Of Users That Are Eligible For Discount](https://leetcode.com/problems/the-number-of-users-that-are-eligible-for-discount/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2205: The Number of Users That Are Eligible for Discount
// https://leetcode.com/problems/the-number-of-users-that-are-eligible-for-discount/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// User purchases: (user_id, amount, date)
	purchases := [][3]string{
		{"1", "120", "2024-01-15"},
		{"2", "50", "2024-02-01"},
		{"1", "80", "2024-03-01"},
		{"3", "200", "2024-01-20"},
	}
	fmt.Println(TheNumberOfUsersThatAreEligibleForDiscount(purchases)) // 2 (users with total > 100)
}

// Time: O(n), Space: O(n)
func TheNumberOfUsersThatAreEligibleForDiscount(purchases [][3]string) int {
  // Membuat map untuk pencarian O(1): key → value
	totals := make(map[int]int)
	for _, p := range purchases {
		id := 0
		for _, c := range p[0] {
			id = id*10 + int(c-'0')
		}
		amount := 0
		for _, c := range p[1] {
			amount = amount*10 + int(c-'0')
		}
		totals[id] += amount
	}

	count := 0
	for _, total := range totals {
		if total >= 100 {
			count++
		}
	}
	return count
}
```
