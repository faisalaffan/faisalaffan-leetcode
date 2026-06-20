# 2990 — Loan Types

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func LoanTypes(loans [][]string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

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
  // Membuat map (HashMap) — pencarian O(1)
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
