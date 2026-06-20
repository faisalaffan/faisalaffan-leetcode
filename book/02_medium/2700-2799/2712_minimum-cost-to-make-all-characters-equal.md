# 2712 — Minimum Cost To Make All Characters Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinimumCostToMakeAllCharactersEqual(s string) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2712: Minimum Cost to Make All Characters Equal
// https://leetcode.com/problems/minimum-cost-to-make-all-characters-equal/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MinimumCostToMakeAllCharactersEqual(s string) int64 {
	n := len(s)

	calc := func(target byte) int64 {
		var cost int64
		flip := 0
		for i := 0; i < n; i++ {
			cur := s[i]
			if flip%2 == 1 {
				if cur == '0' {
					cur = '1'
				} else {
					cur = '0'
				}
			}
			if cur != target {
				flip++
				cost += int64(i + 1)
			}
		}
		return cost
	}

	cost0 := calc('0')
	cost1 := calc('1')
	if cost0 < cost1 {
		return cost0
	}
	return cost1
}

func main() {
	fmt.Println(MinimumCostToMakeAllCharactersEqual("0011"))
	fmt.Println(MinimumCostToMakeAllCharactersEqual("010101"))
}
```
