# 0482 — License Key Formatting

## Deskripsi

**Soal:** [0482. License Key Formatting](https://leetcode.com/problems/license-key-formatting/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func LicenseKeyFormatting(s string, k int) string`

## Solusi Go

```go
package main

// LeetCode #482: License Key Formatting
// https://leetcode.com/problems/license-key-formatting/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func LicenseKeyFormatting(s string, k int) string {
	var result []byte
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		if count == k {
			result = append([]byte{'-'}, result...)
			count = 0
		}
		c := s[i]
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		result = append([]byte{c}, result...)
		count++
	}
	return string(result)
}

func main() {
	fmt.Println(LicenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(LicenseKeyFormatting("2-5g-3-J", 2))
}
```
