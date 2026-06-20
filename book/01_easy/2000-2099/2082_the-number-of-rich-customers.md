# 2082 — The Number Of Rich Customers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data tabel database. Tugasmu adalah menganalisis data tersebut. Karena repo ini Go, query SQL disimulasikan dengan map, slice, dan struct.

**Cara berpikir:** Tentukan SELECT, FROM, JOIN, GROUP BY, ORDER BY. Lalu terjemahkan ke Go.

**Fungsi Solusi:** `func TheNumberOfRichCustomers(transactions [][2]string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2082: The Number of Rich Customers
// https://leetcode.com/problems/the-number-of-rich-customers/
// Difficulty: Easy [Paid] (SQL)

import "fmt"

func main() {
	// (customer_id, amount)
	transactions := [][2]string{{"1", "500"}, {"2", "300"}, {"1", "600"}, {"3", "200"}}
	fmt.Println(TheNumberOfRichCustomers(transactions)) // 1 (only customer 1 has amount > 500)
}

// Time: O(n), Space: O(n)
func TheNumberOfRichCustomers(transactions [][2]string) int {
  // HashMap: O(1) lookup
	customers := make(map[int]bool)
	for _, t := range transactions {
		amount := 0
		for _, c := range t[1] {
			amount = amount*10 + int(c-'0')
		}
		if amount > 500 {
			id := 0
			for _, c := range t[0] {
				id = id*10 + int(c-'0')
			}
			customers[id] = true
		}
	}
	return len(customers)
}
```
