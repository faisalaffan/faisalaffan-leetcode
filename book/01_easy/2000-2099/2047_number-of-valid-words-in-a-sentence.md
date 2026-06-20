# 2047 — Number Of Valid Words In A Sentence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func NumberOfValidWordsInASentence(sentence string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2047: Number of Valid Words in a Sentence
// https://leetcode.com/problems/number-of-valid-words-in-a-sentence/
// Difficulty: Easy

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(NumberOfValidWordsInASentence("cat and  dog"))       // 3
	fmt.Println(NumberOfValidWordsInASentence("!this  1-s b8d!"))     // 0
	fmt.Println(NumberOfValidWordsInASentence("alice and  bob are playing stone-game10")) // 5
}

// Time: O(n), Space: O(1)
func NumberOfValidWordsInASentence(sentence string) int {
	count := 0
	i := 0
	for i < len(sentence) {
		// Skip spaces
		if sentence[i] == ' ' {
			i++
			continue
		}

		// Extract word
		start := i
		for i < len(sentence) && sentence[i] != ' ' {
			i++
		}
		word := sentence[start:i]

		if isValidWord(word) {
			count++
		}
	}
	return count
}

func isValidWord(word string) bool {
	hasHyphen := false
	for i, ch := range word {
		if unicode.IsDigit(ch) {
			return false
		}
		if ch == '-' {
			if hasHyphen || i == 0 || i == len(word)-1 {
				return false
			}
			if !unicode.IsLower(rune(word[i-1])) || !unicode.IsLower(rune(word[i+1])) {
				return false
			}
			hasHyphen = true
		}
		if ch == '!' || ch == '.' || ch == ',' {
			if i != len(word)-1 {
				return false
			}
		}
	}
	return true
}
```
