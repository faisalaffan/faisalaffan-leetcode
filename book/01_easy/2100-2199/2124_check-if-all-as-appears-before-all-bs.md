# 2124 — Check If All As Appears Before All Bs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckIfAllAsAppearsBeforeAllBs(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2124: Check if All A's Appears Before All B's
// https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("aaabbb")) // true
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("abab"))   // false
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("bbb"))    // true
}

// Time: O(n), Space: O(1)
func CheckIfAllAsAppearsBeforeAllBs(s string) bool {
	foundB := false
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			foundB = true
		} else if s[i] == 'a' && foundB {
			return false
		}
	}
	return true
}
```
