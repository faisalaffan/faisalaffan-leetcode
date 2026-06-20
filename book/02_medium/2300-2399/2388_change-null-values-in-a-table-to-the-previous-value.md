# 2388 — Change Null Values In A Table To The Previous Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func intPtr(v int) *int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2388: Change Null Values in a Table to the Previous Value
// https://leetcode.com/problems/change-null-values-in-a-table-to-the-previous-value/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Iterate rows ordered by id, carry forward non-null values.

import "fmt"

type Row struct {
	ID    int
	Value *int // nil represents null
}

func main() {
	rows := []Row{
		{1, intPtr(3)},
		{2, nil},
		{3, nil},
		{4, intPtr(6)},
		{5, nil},
	}
	fillNull(rows)
	for _, r := range rows {
		if r.Value != nil {
			fmt.Printf("%d:%d ", r.ID, *r.Value)
		} else {
			fmt.Printf("%d:nil ", r.ID)
		}
	}
	fmt.Println()
	// Output: 1:3 2:3 3:3 4:6 5:6
}

func intPtr(v int) *int { return &v }

func fillNull(rows []Row) {
	var prev *int
  // Range loop
	for i := range rows {
		if rows[i].Value != nil {
			prev = rows[i].Value
		} else if prev != nil {
			rows[i].Value = intPtr(*prev)
		}
	}
}
```
