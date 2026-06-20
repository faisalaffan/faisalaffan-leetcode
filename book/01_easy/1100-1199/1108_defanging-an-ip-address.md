# 1108 — Defanging An Ip Address

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func defangIPaddr(address string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1108: Defanging an IP Address
// https://leetcode.com/problems/defanging-an-ip-address/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))       // "1[.]1[.]1[.]1"
	fmt.Println(defangIPaddr("255.100.50.0"))  // "255[.]100[.]50[.]0"
}

// LeetCode submission: defangIPaddr
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
```
