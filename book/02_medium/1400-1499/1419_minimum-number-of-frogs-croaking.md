# 1419 — Minimum Number Of Frogs Croaking

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minNumberOfFrogs(croakOfFrogs string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) where n = length of croakOfFrogs  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1419: Minimum Number of Frogs Croaking
// https://leetcode.com/problems/minimum-number-of-frogs-croaking/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minNumberOfFrogs("croakcroak")) // 1

	// Test case 2
	fmt.Println(minNumberOfFrogs("crcoakroak")) // 2

	// Test case 3
	fmt.Println(minNumberOfFrogs("croakcrook")) // -1

	// Test case 4
	fmt.Println(minNumberOfFrogs("croakcroa")) // -1
}

// Time: O(n) where n = length of croakOfFrogs
// Space: O(1)
func minNumberOfFrogs(croakOfFrogs string) int {
	// Track count of each character in the sequence c->r->o->a->k
  // Alokasi slice
	count := make([]int, 5)
	active := 0
	maxFrogs := 0

	for _, ch := range croakOfFrogs {
		switch ch {
		case 'c':
			count[0]++
			active++
			if active > maxFrogs {
				maxFrogs = active
			}
		case 'r':
			if count[0] <= count[1] {
				return -1
			}
			count[1]++
		case 'o':
			if count[1] <= count[2] {
				return -1
			}
			count[2]++
		case 'a':
			if count[2] <= count[3] {
				return -1
			}
			count[3]++
		case 'k':
			if count[3] <= count[4] {
				return -1
			}
			count[4]++
			active--
		default:
			return -1
		}
	}

	// All counts must be equal (complete "croak" sequences)
	if count[0] != count[1] || count[1] != count[2] || count[2] != count[3] || count[3] != count[4] {
		return -1
	}

	return maxFrogs
}
```
