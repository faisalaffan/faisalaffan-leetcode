# 3582 — Generate Tag For Video Caption

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func GenerateTagForVideoCaption(caption string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3582: Generate Tag for Video Caption
// https://leetcode.com/problems/generate-tag-for-video-caption/
// Difficulty: Easy

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(GenerateTagForVideoCaption("Leetcode daily streak achieved"))
	fmt.Println(GenerateTagForVideoCaption("can I Go There"))
	fmt.Println(GenerateTagForVideoCaption("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"))
}

// Time: O(n)
// Space: O(n)
func GenerateTagForVideoCaption(caption string) string {
	words := strings.Fields(caption)
	for i, w := range words {
		if i == 0 {
			words[i] = strings.ToLower(w)
		} else {
			runes := []rune(w)
			for j, r := range runes {
				if j == 0 {
					runes[j] = unicode.ToUpper(r)
				} else {
					runes[j] = unicode.ToLower(r)
				}
			}
			words[i] = string(runes)
		}
	}

	result := "#" + strings.Join(words, "")
	if len(result) > 100 {
		result = result[:100]
	}
	return result
}
```
