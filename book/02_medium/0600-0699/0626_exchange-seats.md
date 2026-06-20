# 0626 — Exchange Seats

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func ExchangeSeats(students [][]interface{}) [][]interface`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // Matriks 2D
	result := make([][]interface{}, n)

	// Build map for easy lookup
  // HashMap: O(1) lookup
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
