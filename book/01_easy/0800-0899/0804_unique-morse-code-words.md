# 0804 — Unique Morse Code Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func uniqueMorseRepresentations(words []string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * m) where n = len(words), m = avg len. Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #804: Unique Morse Code Words
// https://leetcode.com/problems/unique-morse-code-words/
// Difficulty: Easy

import "fmt"

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})) // 2
	fmt.Println(uniqueMorseRepresentations([]string{"a"}))                        // 1
}

// uniqueMorseRepresentations counts unique Morse code transformations of words.
// Time: O(n * m) where n = len(words), m = avg len. Space: O(n).
func uniqueMorseRepresentations(words []string) int {
  // HashMap: O(1) lookup
	set := make(map[string]bool)
	for _, word := range words {
		var code string
		for _, c := range word {
			code += morse[c-'a']
		}
		set[code] = true
	}
	return len(set)
}
```
