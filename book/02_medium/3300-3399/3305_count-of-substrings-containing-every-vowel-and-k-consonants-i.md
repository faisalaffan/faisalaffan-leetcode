# 3305 — Count Of Substrings Containing Every Vowel And K Consonants I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countOfSubstrings(word string, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^2) Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3305: Count of Substrings Containing Every Vowel and K Consonants I
// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-i/
// Difficulty: Medium
// Time: O(n^2) Space: O(1)

import "fmt"

func main() {
	fmt.Println(countOfSubstrings("aeioqq", 1))         // 0
	fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1)) // 3
	fmt.Println(countOfSubstrings("aeiou", 0))          // 1
}

func countOfSubstrings(word string, k int) int64 {
	n := len(word)
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	var ans int64
	for i := 0; i < n; i++ {
  // HashMap: O(1) lookup
		vowelCnt := make(map[byte]int)
		cons := 0
		for j := i; j < n; j++ {
			if isVowel(word[j]) {
				vowelCnt[word[j]]++
			} else {
				cons++
				if cons > k {
					break
				}
			}
			if cons == k && len(vowelCnt) == 5 {
				ans++
			}
		}
	}
	return ans
}
```
