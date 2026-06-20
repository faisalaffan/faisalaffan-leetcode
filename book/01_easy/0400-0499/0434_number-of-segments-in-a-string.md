# 0434 — Number Of Segments In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func NumberOfSegmentsInAString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #434: Number of Segments in a String
// https://leetcode.com/problems/number-of-segments-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func NumberOfSegmentsInAString(s string) int {
	count := 0
	inSegment := false
	for _, c := range s {
		if c != ' ' && !inSegment {
			count++
			inSegment = true
		} else if c == ' ' {
			inSegment = false
		}
	}
	return count
}

func main() {
	fmt.Println(NumberOfSegmentsInAString("Hello, my name is John"))
	fmt.Println(NumberOfSegmentsInAString("Hello"))
	fmt.Println(NumberOfSegmentsInAString(""))
}
```
