# 3773 — Maximum Number Of Equal Length Runs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func maximumNumberOfEqualLengthRuns(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3773: Maximum Number of Equal Length Runs
// https://leetcode.com/problems/maximum-number-of-equal-length-runs/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func maximumNumberOfEqualLengthRuns(s string) int {
  // HashMap: O(1) lookup
	cnt := make(map[int]int)
	maxCount := 0
	n := len(s)
	i := 0
	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		runLen := j - i
		cnt[runLen]++
		if cnt[runLen] > maxCount {
			maxCount = cnt[runLen]
		}
		i = j
	}
	return maxCount
}

func main() {
	fmt.Println(maximumNumberOfEqualLengthRuns("hello"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aaabaaa"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aabbcc"))
}
```
