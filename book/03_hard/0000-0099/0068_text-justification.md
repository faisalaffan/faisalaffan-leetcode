# 0068 — Text Justification

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func fullJustify(words []string, maxWidth int) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #68: Text Justification
// https://leetcode.com/problems/text-justification/
// Difficulty: Hard

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("68. Text Justification")
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	result := fullJustify(words, 16)
	fmt.Printf("maxWidth=16 -> %q\n", result)
	for _, line := range result {
		fmt.Printf("  %q (len=%d)\n", line, len(line))
	}

	words2 := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	result2 := fullJustify(words2, 16)
	fmt.Printf("\nmaxWidth=16 -> %q\n", result2)
	for _, line := range result2 {
		fmt.Printf("  %q (len=%d)\n", line, len(line))
	}

	words3 := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	result3 := fullJustify(words3, 20)
	fmt.Printf("\nmaxWidth=20 -> %q\n", result3)
	for _, line := range result3 {
		fmt.Printf("  %q (len=%d)\n", line, len(line))
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	i := 0
	for i < len(words) {
		j := i + 1
		lineLen := len(words[i])
		for j < len(words) && lineLen+1+len(words[j]) <= maxWidth {
			lineLen += 1 + len(words[j])
			j++
		}

		line := ""
		wordCount := j - i
		spaceSlots := wordCount - 1

		if j == len(words) || wordCount == 1 {
			// left-justified (last line or single word)
			line = strings.Join(words[i:j], " ")
			line += strings.Repeat(" ", maxWidth-len(line))
		} else {
			totalSpaces := maxWidth - (lineLen - spaceSlots) // total spaces needed
			baseSpaces := totalSpaces / spaceSlots
			extraSpaces := totalSpaces % spaceSlots

			var sb strings.Builder
			for k := i; k < j; k++ {
				sb.WriteString(words[k])
				if k < j-1 {
					spaces := baseSpaces
					if k-i < extraSpaces {
						spaces++
					}
					sb.WriteString(strings.Repeat(" ", spaces))
				}
			}
			line = sb.String()
		}
		result = append(result, line)
		i = j
	}
	return result
}
```
