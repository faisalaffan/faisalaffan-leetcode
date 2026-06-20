# 0748 — Shortest Completing Word

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func shortestCompletingWord(licensePlate string, words []string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m) where n = len(words), m = avg word length. Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #748: Shortest Completing Word
// https://leetcode.com/problems/shortest-completing-word/
// Difficulty: Easy

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"})) // "steps"
	fmt.Println(shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))     // "pest"
}

// shortestCompletingWord finds the shortest word that contains all letters in licensePlate.
// Time: O(n * m) where n = len(words), m = avg word length. Space: O(1).
func shortestCompletingWord(licensePlate string, words []string) string {
	targetCount := [26]int{}
	for _, c := range strings.ToLower(licensePlate) {
		if c >= 'a' && c <= 'z' {
			targetCount[c-'a']++
		}
	}

	result := ""
	for _, word := range words {
		if result != "" && len(word) >= len(result) {
			continue
		}
		wordCount := [26]int{}
		for _, c := range word {
			wordCount[c-'a']++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if wordCount[i] < targetCount[i] {
				ok = false
				break
			}
		}
		if ok {
			result = word
		}
	}
	return result
}
```
