# 0246 — Strobogrammatic Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func IsStrobogrammatic(num string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #246: Strobogrammatic Number
// https://leetcode.com/problems/strobogrammatic-number/
// Difficulty: Easy [Paid]

import "fmt"

// Time: O(n) | Space: O(1)
func IsStrobogrammatic(num string) bool {
	pairs := map[byte]byte{'0': '0', '1': '1', '6': '9', '8': '8', '9': '6'}
	i, j := 0, len(num)-1
	for i <= j {
		if v, ok := pairs[num[i]]; !ok || v != num[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(IsStrobogrammatic("69"))
	fmt.Println(IsStrobogrammatic("88"))
	fmt.Println(IsStrobogrammatic("962"))
}
```
