# 0831 — Masking Personal Information

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaskingPersonalInformation(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #831: Masking Personal Information
// https://leetcode.com/problems/masking-personal-information/
// Difficulty: Medium

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(MaskingPersonalInformation("LeetCode@LeetCode.com"))
	fmt.Println(MaskingPersonalInformation("AB@qq.com"))
	fmt.Println(MaskingPersonalInformation("1(234)567-890"))
	fmt.Println(MaskingPersonalInformation("86-(10)12345678"))
}

// Time: O(n) | Space: O(n)
func MaskingPersonalInformation(s string) string {
	at := strings.IndexByte(s, '@')
	if at != -1 {
		s = strings.ToLower(s)
		return string(s[0]) + "*****" + string(s[at-1]) + s[at:]
	}

	var digits strings.Builder
	for _, c := range s {
		if unicode.IsDigit(c) {
			digits.WriteRune(c)
		}
	}
	d := digits.String()
	local := d[len(d)-4:]
	masked := "***-***-" + local

	if len(d) == 10 {
		return masked
	}
	return "+" + strings.Repeat("*", len(d)-10) + "-" + masked
}
```
