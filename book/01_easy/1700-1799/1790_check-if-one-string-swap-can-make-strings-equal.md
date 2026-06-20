# 1790 — Check If One String Swap Can Make Strings Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func AreAlmostEqual(s1 string, s2 string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1790: Check if One String Swap Can Make Strings Equal
// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func AreAlmostEqual(s1 string, s2 string) bool {
	var diff []int
  // Linear scan O(n)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}
	if len(diff) == 0 {
		return true
	}
	if len(diff) != 2 {
		return false
	}
	return s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]]
}

func main() {
	fmt.Println(AreAlmostEqual("bank", "kanb"))
	fmt.Println(AreAlmostEqual("attack", "defend"))
	fmt.Println(AreAlmostEqual("kelb", "kelb"))
}
```
