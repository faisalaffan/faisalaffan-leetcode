# 2042 — Check If Numbers Are Ascending In A Sentence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckIfNumbersAreAscendingInASentence(s string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2042: Check if Numbers Are Ascending in a Sentence
// https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/
// Difficulty: Easy

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(CheckIfNumbersAreAscendingInASentence("1 box has 3 blue 4 red 6 green and 12 yellow marbles")) // true
	fmt.Println(CheckIfNumbersAreAscendingInASentence("hello world 5 x 5"))                                     // false
	fmt.Println(CheckIfNumbersAreAscendingInASentence("sunset is at 7 11 pm overnight"))                        // false
}

// Time: O(n), Space: O(1)
func CheckIfNumbersAreAscendingInASentence(s string) bool {
	prev := 0
	num := 0
	hasNum := false

  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			num = num*10 + int(s[i]-'0')
			hasNum = true
		} else {
			if hasNum {
				if num <= prev {
					return false
				}
				prev = num
				num = 0
				hasNum = false
			}
		}
	}
	if hasNum && num <= prev {
		return false
	}
	return true
}
```
