# 0205 — Isomorphic Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func IsIsomorphic(s string, t string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1) (fixed ASCII chars)


## 💻 Solusi Go

```go
package main

// LeetCode #205: Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (fixed ASCII chars)
func IsIsomorphic(s string, t string) bool {
  // Alokasi slice
	m1 := make([]int, 256)
  // Alokasi slice
	m2 := make([]int, 256)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

func main() {
	fmt.Println(IsIsomorphic("egg", "add"))
	fmt.Println(IsIsomorphic("foo", "bar"))
	fmt.Println(IsIsomorphic("paper", "title"))
}
```
