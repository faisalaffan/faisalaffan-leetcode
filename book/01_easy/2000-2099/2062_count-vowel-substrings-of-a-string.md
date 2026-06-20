# 2062 — Count Vowel Substrings Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountVowelSubstringsOfAString(word string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^2), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2062: Count Vowel Substrings of a String
// https://leetcode.com/problems/count-vowel-substrings-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountVowelSubstringsOfAString("aeiouu"))    // 2
	fmt.Println(CountVowelSubstringsOfAString("unicornarihan")) // 0
	fmt.Println(CountVowelSubstringsOfAString("cuaieuouac"))    // 7
}

// Time: O(n^2), Space: O(1)
func CountVowelSubstringsOfAString(word string) int {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	count := 0
  // Linear scan O(n)
	for i := 0; i < len(word); i++ {
  // HashMap: O(1) lookup
		vowelSet := make(map[byte]bool)
		for j := i; j < len(word); j++ {
			if !isVowel(word[j]) {
				break
			}
			vowelSet[word[j]] = true
			if len(vowelSet) == 5 {
				count++
			}
		}
	}
	return count
}
```
