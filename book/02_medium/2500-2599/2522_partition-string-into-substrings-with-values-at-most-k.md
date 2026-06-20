# 2522 — Partition String Into Substrings With Values At Most K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumPartition(s string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2522: Partition String Into Substrings With Values at Most K
// https://leetcode.com/problems/partition-string-into-substrings-with-values-at-most-k/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: extend substring while value <= k, then start new partition.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumPartition("165462", 60)) // 4 (16|54|6|2)
	fmt.Println(minimumPartition("238182", 5))  // -1
}

func minimumPartition(s string, k int) int {
	ans := 1
	cur := 0
	for _, ch := range s {
		d := int(ch - '0')
		if d > k {
			return -1
		}
		if cur > math.MaxInt32/10 || cur*10+d > k {
			ans++
			cur = d
		} else {
			cur = cur*10 + d
		}
	}
	return ans
}
```
