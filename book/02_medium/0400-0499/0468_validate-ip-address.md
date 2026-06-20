# 0468 — Validate Ip Address

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func ValidateIpAddress(queryIP string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #468: Validate IP Address
// https://leetcode.com/problems/validate-ip-address/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(ValidateIpAddress("172.16.254.1"))
	fmt.Println(ValidateIpAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(ValidateIpAddress("256.256.256.256"))
}

func ValidateIpAddress(queryIP string) string {
	if isIPv4(queryIP) {
		return "IPv4"
	}
	if isIPv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func isIPv4(s string) bool {
	parts := strings.Split(s, ".")
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if len(p) == 0 || (len(p) > 1 && p[0] == '0') {
			return false
		}
		num, err := strconv.Atoi(p)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func isIPv6(s string) bool {
	parts := strings.Split(s, ":")
	if len(parts) != 8 {
		return false
	}
	for _, p := range parts {
		if len(p) == 0 || len(p) > 4 {
			return false
		}
		for _, c := range p {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
				return false
			}
		}
	}
	return true
}
```
