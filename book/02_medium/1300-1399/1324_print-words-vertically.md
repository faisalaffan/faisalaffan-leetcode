# 1324 — Print Words Vertically

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func printVertically(s string) []string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m*n) where m = max word length, n = number of words  |  **Ruang:** O(m*n) for the result


## 💻 Solusi Go

```go
package main

// LeetCode #1324: Print Words Vertically
// https://leetcode.com/problems/print-words-vertically/
// Difficulty: Medium

import "fmt"
import "strings"

func main() {
	// Test case 1
	fmt.Println(printVertically("HOW ARE YOU"))
	// ["HAY","ORO","WEU"]

	// Test case 2
	fmt.Println(printVertically("TO BE OR NOT TO BE"))
	// ["TBONTB","OEROOE","   T"]

	// Test case 3
	fmt.Println(printVertically("CONTEST IS COMING"))
	// ["CIC","OSO","N M","T I","E N","S G","T"]
}

// Time: O(m*n) where m = max word length, n = number of words
// Space: O(m*n) for the result
func printVertically(s string) []string {
	words := strings.Fields(s)
	maxLen := 0
	for _, w := range words {
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	result := make([]string, maxLen)
	for i := 0; i < maxLen; i++ {
		var sb strings.Builder
		for _, w := range words {
			if i < len(w) {
				sb.WriteByte(w[i])
			} else {
				sb.WriteByte(' ')
			}
		}
		// Trim trailing spaces
		result[i] = strings.TrimRight(sb.String(), " ")
	}

	return result
}
```
