# 3146 — Permutation Difference Between Two Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func PermutationDifferenceBetweenTwoStrings(s string, t string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3146: Permutation Difference between Two Strings
// https://leetcode.com/problems/permutation-difference-between-two-strings/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findPermutationDifference
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abc", "bac")) // 2
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abcde", "edcba")) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: findPermutationDifference
func PermutationDifferenceBetweenTwoStrings(s string, t string) int {
  // HashMap: O(1) lookup
	pos := make(map[byte]int)
  // Linear scan O(n)
	for i := 0; i < len(t); i++ {
		pos[t[i]] = i
	}
	diff := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		d := i - pos[s[i]]
		if d < 0 {
			d = -d
		}
		diff += d
	}
	return diff
}
```
