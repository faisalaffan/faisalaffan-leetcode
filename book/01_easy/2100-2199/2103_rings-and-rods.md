# 2103 — Rings And Rods

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func RingsAndRods(rings string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2103: Rings and Rods
// https://leetcode.com/problems/rings-and-rods/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RingsAndRods("B0B6G0R6R0R6G9"))    // 1
	fmt.Println(RingsAndRods("B0R0G0R9R0B0G0"))    // 1
	fmt.Println(RingsAndRods("G4"))                 // 0
}

// Time: O(n), Space: O(1)
func RingsAndRods(rings string) int {
  // Alokasi slice
	rods := make([]int, 10)
  // Linear scan O(n)
	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		rod := rings[i+1] - '0'
		switch color {
		case 'R':
			rods[rod] |= 1
		case 'G':
			rods[rod] |= 2
		case 'B':
			rods[rod] |= 4
		}
	}

	count := 0
	for _, v := range rods {
		if v == 7 { // R|G|B = 1|2|4 = 7
			count++
		}
	}
	return count
}
```
