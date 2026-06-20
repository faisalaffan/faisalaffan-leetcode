# 2125 — Number Of Laser Beams In A Bank

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func numberOfBeams(bank []string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2125: Number of Laser Beams in a Bank
// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/
// Difficulty: Medium
// Time: O(m*n) | Space: O(1)

import "fmt"

func numberOfBeams(bank []string) int {
	prevCount := 0
	result := 0

	for _, row := range bank {
		count := 0
		for _, c := range row {
			if c == '1' {
				count++
			}
		}
		if count > 0 {
			result += prevCount * count
			prevCount = count
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
	// Expected: 8

	// Test case 2
	fmt.Println("Test 2:", numberOfBeams([]string{"000", "111", "000"}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", numberOfBeams([]string{"101", "010", "101"}))
	// Expected: 4
}
```
