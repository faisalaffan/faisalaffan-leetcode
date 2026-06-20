# 2405 — Optimal Partition Of String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func partitionString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2405: Optimal Partition of String
// https://leetcode.com/problems/optimal-partition-of-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Greedy: start new partition when duplicate char found.

import "fmt"

func main() {
	fmt.Println(partitionString("abacaba")) // 4
	fmt.Println(partitionString("ssssss"))  // 6
}

func partitionString(s string) int {
	ans := 1
  // HashMap: O(1) lookup
	seen := make(map[byte]bool)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if seen[s[i]] {
			ans++
			seen = make(map[byte]bool)
		}
		seen[s[i]] = true
	}
	return ans
}
```
