# 1271 — Hexspeak

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func toHexspeak(num string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(log n)  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1271: Hexspeak
// https://leetcode.com/problems/hexspeak/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(log n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(toHexspeak("257"))  // "IOI"
	fmt.Println(toHexspeak("3"))    // "ERROR"
	fmt.Println(toHexspeak("619"))  // "ERROR" (619=26B, B not allowed)
}

// LeetCode submission: toHexspeak
func toHexspeak(num string) string {
	n, _ := strconv.Atoi(num)
	hex := strconv.FormatInt(int64(n), 16)
	replacer := map[byte]byte{
		'0': 'O',
		'1': 'I',
	}
	ans := make([]byte, len(hex))
  // Range loop
	for i := range hex {
		if r, ok := replacer[hex[i]]; ok {
			ans[i] = r
		} else if hex[i] >= 'a' && hex[i] <= 'f' {
			ans[i] = hex[i] - 'a' + 'A'
		} else {
			return "ERROR"
		}
	}
	return string(ans)
}
```
