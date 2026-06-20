# 0860 — Lemonade Change

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func lemonadeChange(bills []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #860: Lemonade Change
// https://leetcode.com/problems/lemonade-change/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20})) // true
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20})) // false
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 5, 20, 5, 10, 5, 20})) // true
}

// lemonadeChange checks if we can provide correct change for each customer.
// Time: O(n). Space: O(1).
func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			fives++
		case 10:
			if fives == 0 {
				return false
			}
			fives--
			tens++
		case 20:
			if tens > 0 && fives > 0 {
				tens--
				fives--
			} else if fives >= 3 {
				fives -= 3
			} else {
				return false
			}
		}
	}
	return true
}
```
