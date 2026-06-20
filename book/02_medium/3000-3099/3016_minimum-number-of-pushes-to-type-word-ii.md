# 3016 — Minimum Number Of Pushes To Type Word Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimumPushes(word string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3016: Minimum Number of Pushes to Type Word II
// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumPushes("abcde"))
	fmt.Println(minimumPushes("xyzxyzxyzxyz"))
	fmt.Println(minimumPushes("aabbccddeeffgghhiiiiii"))
}

func minimumPushes(word string) int {
  // Alokasi slice
	cnt := make([]int, 26)
	for _, ch := range word {
		cnt[ch-'a']++
	}
  // Custom sort
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})
	ans := 0
	for i, c := range cnt {
		ans += c * (i/8 + 1)
	}
	return ans
}
```
