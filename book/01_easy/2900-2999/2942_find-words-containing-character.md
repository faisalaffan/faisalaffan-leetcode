# 2942 — Find Words Containing Character

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func FindWordsContainingCharacter(words []string, x byte) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m) where m is max word length  |  **Ruang:** O(1) excluding output


## 💻 Solusi Go

```go
package main

// LeetCode #2942: Find Words Containing Character
// https://leetcode.com/problems/find-words-containing-character/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findWordsContaining
	fmt.Println(FindWordsContainingCharacter([]string{"leet", "code"}, 'e')) // [0, 1]
	fmt.Println(FindWordsContainingCharacter([]string{"abc", "bcd", "aaaa", "cbc"}, 'a')) // [0, 2]
	fmt.Println(FindWordsContainingCharacter([]string{"abc", "bcd", "aaaa", "cbc"}, 'z')) // []
}

// Time: O(n * m) where m is max word length | Space: O(1) excluding output
// LeetCode submission name: findWordsContaining
func FindWordsContainingCharacter(words []string, x byte) []int {
	result := []int{}
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			if word[j] == x {
				result = append(result, i)
				break
			}
		}
	}
	return result
}
```
