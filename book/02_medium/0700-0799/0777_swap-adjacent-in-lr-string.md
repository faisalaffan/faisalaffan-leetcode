# 0777 — Swap Adjacent In Lr String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func canTransform(start string, end string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #777: Swap Adjacent in LR String
// https://leetcode.com/problems/swap-adjacent-in-lr-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
	fmt.Println(canTransform("X", "L"))
}

func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}

	n := len(start)
	i, j := 0, 0

	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i++
		}
		for j < n && end[j] == 'X' {
			j++
		}

		if i == n && j == n {
			return true
		}
		if i == n || j == n {
			return false
		}
		if start[i] != end[j] {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}

		i++
		j++
	}

	return true
}
```
