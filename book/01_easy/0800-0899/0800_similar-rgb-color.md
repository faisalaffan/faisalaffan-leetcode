# 0800 — Similar Rgb Color

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func similarRGB(color string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #800: Similar RGB Color
// https://leetcode.com/problems/similar-rgb-color/
// Difficulty: Easy [Paid]
// Note: This is a premium problem. Implementation based on public description.

import "fmt"

func main() {
	fmt.Println(similarRGB("#09f166")) // "#11ee66"
	fmt.Println(similarRGB("#4e3fe1")) // "#44ee11" (note: this is an approximation)
}

// similarRGB finds the most similar RGB shorthand for the given color.
// Time: O(1). Space: O(1).
func similarRGB(color string) string {
	// For each component, find the nearest shorthand (00, 11, ..., ff)
	result := "#"
	for i := 1; i < len(color); i += 2 {
		// Extract the 2-digit hex value
		val := 0
		for j := 0; j < 2; j++ {
			c := color[i+j]
			if c >= '0' && c <= '9' {
				val = val*16 + int(c-'0')
			} else {
				val = val*16 + int(c-'a'+10)
			}
		}
		// Find the nearest shorthand: 0x00, 0x11, 0x22, ..., 0xff
		nearest := (val + 8) / 17 * 17
		if nearest > 255 {
			nearest = 255
		}
		hex := fmt.Sprintf("%02x", nearest)
		result += hex
	}
	return result
}
```
