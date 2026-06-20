# 0626 — Exchange Seats

## Deskripsi

**Soal:** [0626. Exchange Seats](https://leetcode.com/problems/exchange-seats/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #626: Exchange Seats
// https://leetcode.com/problems/exchange-seats/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	// Students: {id, name}
	students := [][]interface{}{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
		{4, "David"},
	}
	result := ExchangeSeats(students)
	for _, r := range result {
		fmt.Printf("id=%d, name=%s\n", r[0].(int), r[1].(string))
	}
}

func ExchangeSeats(students [][]interface{}) [][]interface{} {
	n := len(students)
  // Membuat slice 2D untuk DP/tabel
	result := make([][]interface{}, n)

	// Build map for easy lookup
  // Membuat map untuk pencarian O(1): key → value
	studentMap := make(map[int]string)
	for _, student := range students {
		id := student[0].(int)
		name := student[1].(string)
		studentMap[id] = name
	}

	for id := 1; id <= n; id++ {
		if id%2 == 1 {
			if id+1 <= n {
				result[id-1] = []interface{}{id + 1, studentMap[id+1]}
			} else {
				result[id-1] = []interface{}{id, studentMap[id]}
			}
		} else {
			result[id-1] = []interface{}{id - 1, studentMap[id-1]}
		}
	}

	return result
}
```
