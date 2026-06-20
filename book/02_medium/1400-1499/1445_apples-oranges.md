# 1445 — Apples Oranges

## Deskripsi

**Soal:** [1445. Apples Oranges](https://leetcode.com/problems/apples-oranges/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1445: Apples & Oranges
// https://leetcode.com/problems/apples-oranges/
// Difficulty: Medium

import "fmt"

func main() {
	// SQL problem - simulating in Go
	result := applesOranges([]struct {
		saleDate string
		fruit    string
		soldNum  int
	}{
		{"2020-05-01", "apples", 10},
		{"2020-05-01", "oranges", 8},
		{"2020-05-02", "apples", 15},
		{"2020-05-02", "oranges", 15},
		{"2020-05-03", "apples", 20},
		{"2020-05-03", "oranges", 0},
		{"2020-05-04", "apples", 15},
		{"2020-05-04", "oranges", 16},
	})
	for _, r := range result {
		fmt.Printf("%s %d\n", r.date, r.diff)
	}
}

type diffResult struct {
	date string
	diff int
}

// Time: O(n log n) for sorting
// Space: O(n)
func applesOranges(sales []struct {
	saleDate string
	fruit    string
	soldNum  int
}) []diffResult {
	// Group by date
  // Membuat map untuk pencarian O(1): key → value
	apples := make(map[string]int)
  // Membuat map untuk pencarian O(1): key → value
	oranges := make(map[string]int)

  // Membuat map untuk pencarian O(1): key → value
	dateSet := make(map[string]bool)
	for _, s := range sales {
		dateSet[s.saleDate] = true
		if s.fruit == "apples" {
			apples[s.saleDate] += s.soldNum
		} else {
			oranges[s.saleDate] += s.soldNum
		}
	}

	// Sort dates
  // Membuat slice untuk menyimpan hasil
	dates := make([]string, 0, len(dateSet))
	for d := range dateSet {
		dates = append(dates, d)
	}

	var result []diffResult
	for _, d := range dates {
		diff := apples[d] - oranges[d]
		result = append(result, diffResult{d, diff})
	}

	return result
}
```
