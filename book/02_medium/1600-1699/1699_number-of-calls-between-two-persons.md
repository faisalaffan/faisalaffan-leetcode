# 1699 — Number Of Calls Between Two Persons

## Deskripsi

**Soal:** [1699. Number Of Calls Between Two Persons](https://leetcode.com/problems/number-of-calls-between-two-persons/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfCalls(records []CallRecord) [][3]int`

## Solusi Go

```go
package main

// LeetCode #1699: Number of Calls Between Two Persons
// https://leetcode.com/problems/number-of-calls-between-two-persons/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n), Space: O(n)

import "fmt"

type CallRecord struct {
	FromID int
	ToID   int
	Dur    int
}

func numberOfCalls(records []CallRecord) [][3]int {
  // Membuat map untuk pencarian O(1): key → value
	callMap := make(map[[2]int]int) // [min,max] -> total duration

	for _, r := range records {
		a, b := r.FromID, r.ToID
		if a > b {
			a, b = b, a
		}
		key := [2]int{a, b}
		callMap[key] += r.Dur
	}

  // Membuat slice untuk menyimpan hasil
	result := make([][3]int, 0, len(callMap))
	for key, dur := range callMap {
		result = append(result, [3]int{key[0], key[1], dur})
	}
	return result
}

func main() {
	records := []CallRecord{
		{1, 2, 10},
		{2, 1, 20},
		{1, 3, 30},
	}
	result := numberOfCalls(records)
	for _, r := range result {
		fmt.Printf("(%d,%d): %d\n", r[0], r[1], r[2])
	}
}
```
