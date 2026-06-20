# 2990 — Loan Types

## Deskripsi

**Soal:** [2990. Loan Types](https://leetcode.com/problems/loan-types/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2990: Loan Types
// https://leetcode.com/problems/loan-types/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: find users who have both 'Refinance' and 'Mortgage' loans.

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: loanTypes
	fmt.Println(LoanTypes([][]string{
		{"101", "Mortgage"},
		{"101", "Refinance"},
		{"102", "Mortgage"},
		{"103", "Refinance"},
		{"103", "Student"},
	})) // [101]

	fmt.Println(LoanTypes([][]string{
		{"1", "Refinance"},
		{"1", "Mortgage"},
		{"2", "Refinance"},
		{"2", "Student"},
		{"3", "Mortgage"},
	})) // [1]
}

// Time: O(n log n) | Space: O(n)
// LeetCode submission name: loanTypes
func LoanTypes(loans [][]string) []string {
  // Membuat map untuk pencarian O(1): key → value
	userLoans := make(map[string]map[string]bool)
	for _, row := range loans {
		userID := row[0]
		loanType := row[1]
		if userLoans[userID] == nil {
			userLoans[userID] = make(map[string]bool)
		}
		userLoans[userID][loanType] = true
	}

	result := []string{}
	for userID, types := range userLoans {
		if types["Refinance"] && types["Mortgage"] {
			result = append(result, userID)
		}
	}
	sort.Strings(result)
	return result
}
```
