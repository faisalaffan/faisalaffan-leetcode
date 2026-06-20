# 2752 — Customers With Maximum Number Of Transactions On Consecutive Days

## Deskripsi

**Soal:** [2752. Customers With Maximum Number Of Transactions On Consecutive Days](https://leetcode.com/problems/customers-with-maximum-number-of-transactions-on-consecutive-days/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func customersWithMaxConsecutiveDays(transactions [][]int) []int`

> **Ide Kunci:** Group transactions by customer, sort by day (dedup), find longest

## Solusi Go

```go
package main

// LeetCode #2752: Customers with Maximum Number of Transactions on Consecutive Days
// https://leetcode.com/problems/customers-with-maximum-number-of-transactions-on-consecutive-days/
// Difficulty: Hard [Paid]
//
// Approach: Group transactions by customer, sort by day (dedup), find longest
// consecutive streak. Return customers with the maximum streak, sorted by ID.

import (
	"fmt"
	"sort"
)

type transaction struct {
	customerID int
	day        int
}

func customersWithMaxConsecutiveDays(transactions [][]int) []int {
  // Membuat map untuk pencarian O(1): key → value
	custDays := make(map[int][]int)
	for _, t := range transactions {
		custID, day := t[0], t[1]
		custDays[custID] = append(custDays[custID], day)
	}

	maxStreak := 0
  // Membuat map untuk pencarian O(1): key → value
	custMaxStreak := make(map[int]int)

	for cid, days := range custDays {
		sort.Ints(days)
  // Membuat slice untuk menyimpan hasil
		uniq := make([]int, 0, len(days))
		for i, d := range days {
			if i == 0 || d != days[i-1] {
				uniq = append(uniq, d)
			}
		}

		if len(uniq) == 0 {
			custMaxStreak[cid] = 0
			continue
		}

		streak := 1
		cur := 1
		for i := 1; i < len(uniq); i++ {
			if uniq[i] == uniq[i-1]+1 {
				cur++
				if cur > streak {
					streak = cur
				}
			} else {
				cur = 1
			}
		}

		if streak > maxStreak {
			maxStreak = streak
		}
		custMaxStreak[cid] = streak
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)
	for cid, streak := range custMaxStreak {
		if streak == maxStreak {
			result = append(result, cid)
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	// Example
	fmt.Println(customersWithMaxConsecutiveDays([][]int{
		{1, 1}, {1, 2}, {2, 1}, {2, 2}, {3, 5}, {3, 6}, {3, 7},
	}))
	// Single customer, consecutive
	fmt.Println(customersWithMaxConsecutiveDays([][]int{{1, 1}, {1, 2}, {1, 3}}))
	// No consecutive days
	fmt.Println(customersWithMaxConsecutiveDays([][]int{{1, 1}, {1, 3}, {1, 5}}))
	// Multiple customers same streak
	fmt.Println(customersWithMaxConsecutiveDays([][]int{{1, 1}, {2, 1}}))
	// Empty
	fmt.Println(customersWithMaxConsecutiveDays([][]int{}))
}
```
